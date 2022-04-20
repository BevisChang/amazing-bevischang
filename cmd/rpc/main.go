package main

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	etcd "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"

	"github.com/AmazingTalker/bevis-chang/pkg/dao"
	"github.com/AmazingTalker/bevis-chang/pkg/pb"
	"github.com/AmazingTalker/bevis-chang/pkg/rpc"
	"github.com/AmazingTalker/go-rpc-kit/cachekit"
	"github.com/AmazingTalker/go-rpc-kit/configkit"
	"github.com/AmazingTalker/go-rpc-kit/envkit"
	"github.com/AmazingTalker/go-rpc-kit/flagkit"
	"github.com/AmazingTalker/go-rpc-kit/logkit"
	"github.com/AmazingTalker/go-rpc-kit/metrickit"
	"github.com/AmazingTalker/go-rpc-kit/migrationkit"
	"github.com/AmazingTalker/go-rpc-kit/monitorkit"
	"github.com/AmazingTalker/go-rpc-kit/mysqlkit"
	"github.com/AmazingTalker/go-rpc-kit/rediskit"
	"github.com/AmazingTalker/go-rpc-kit/validatorkit"
)

const (
	projName   = "bevis-chang"
	configRoot = "/configs/envs/development"
)

type ServiceLauncher struct {
	Run    func() error
	Labels []string
}

func main() {

	// get env
	if err := flagkit.Parse(); err != nil {
		panic(err)
	}

	// register env first, it will benefit ctx, logkit, and metric
	envkit.Register(envkit.Config{
		EnvNamespace: env.EnvConfig.Namespace,
		PodName:      env.EnvConfig.PodName,
		ServiceName:  env.EnvConfig.ServiceName,
	})

	// init context
	ctx := context.Background()

	// init logger
	logkit.RegisterAmazingLogger(&logkit.Config{
		Logger:      logkit.LoggerZap,
		Level:       env.LoggerConfig.Level,
		Development: env.LoggerConfig.Development,
		IntegrationAirbrake: &logkit.IntegrationAirbrake{
			ProjectID:  env.LoggerConfig.Airbrake.ProjectID,
			ProjectKey: env.LoggerConfig.Airbrake.ProjectKey,
			Env:        envkit.EnvNamespace(),
		},
	})
	defer logkit.Flush()

	// init etcd
	logkit.Info(ctx, "init etcd", logkit.Payload{
		"addrs":              env.EtcdConfig.Addrs,
		"dialTimeoutSeconds": env.EtcdConfig.DialTimeoutSeconds,
	})

	etcdCli, err := etcd.New(etcd.Config{
		Endpoints:   env.EtcdConfig.Addrs,
		DialTimeout: time.Second * time.Duration(env.EtcdConfig.DialTimeoutSeconds),
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		logkit.Fatal(ctx, "init etcd failed", logkit.Payload{"err": err})
	}
	defer etcdCli.Close()

	// publishing the configs to etcd in development env
	if envkit.Namespace() == envkit.EnvDevelopment {
		publisher := configkit.NewPublisher(etcdCli, configRoot, configkit.RenderRoot(envkit.EnvDevelopment))
		if err := publisher.Publish(ctx); err != nil {
			logkit.Fatal(ctx, "publisher.Publish failed", logkit.Payload{"err": err})
		}
	}

	// init dynamic config watcher
	logkit.Info(ctx, "init config watcher", logkit.Payload{
		"projectName": projName,
		"env":         envkit.Namespace(),
	})

	if err := configkit.LaunchWatcher(ctx, configkit.Params{
		ProjectName: projName,
		Client:      etcdCli,
		Env:         envkit.Namespace(),
	}); err != nil {
		logkit.Fatal(ctx, "configkit.LaunchWatcher failed", logkit.Payload{"err": err})
	}

	// init metric
	logkit.Info(ctx, "init metric", logkit.Payload{
		"url":            env.MetricConfig.URL,
		"refleshSeconds": env.MetricConfig.RefleshSeconds,
	})

	metrickit.Register(metrickit.Config{
		Exporter:      metrickit.ExporterNewRelic,
		URL:           env.MetricConfig.URL,
		APIKey:        env.MetricConfig.APIKey,
		RefleshPeriod: time.Duration(time.Duration(env.RefleshSeconds) * time.Second),
	})

	// init monitoring
	logkit.Info(ctx, "init monitor", logkit.Payload{
		"period-in-secs": env.MonitorConfig.PeriodSecs,
	})

	monitorkit.Register(monitorkit.Config{
		Metric:                   metrickit.New("monitor"),
		RuntimeCollectorInterval: time.Duration(time.Duration(env.MonitorConfig.PeriodSecs) * time.Second),
	})
	monitorkit.Run()
	defer monitorkit.GracefulStop()

	// init redis
	logkit.Info(ctx, "init redis", logkit.Payload{"addrs": env.RedisConfig.Addrs})
	ring, err := rediskit.NewRedisRing(env.RedisConfig.Addrs)
	if err != nil {
		logkit.Fatal(ctx, "init redis failed", logkit.Payload{"err": err})
	}

	// init cache
	logkit.Info(ctx, "init cache", logkit.Payload{"size": env.LocalCacheConfig.Size})
	cacheSrv := cachekit.NewCache(
		cachekit.NewSharedCache(ring),
		cachekit.NewLocalCache(env.LocalCacheConfig.Size),
	)

	// init db conn
	db, err := mysqlkit.NewMySqlConn(mysqlkit.MySqlConnOpt{Config: &env.MysqlConnConfig})
	if err != nil {
		logkit.Fatal(ctx, "init db conn failed", logkit.Payload{"err": err})
	}

	// https://gorm.io/docs/generic_interface.html
	sqlDB, err := db.DB()
	if err != nil {
		logkit.Fatal(ctx, "invalid db", logkit.Payload{"err": err})
	}
	defer sqlDB.Close()

	// db migration check
	logkit.Infof(ctx, "start migration")
	migrationKit := migrationkit.NewGooseMigrationKit(migrationkit.GooseMysqlDriver, migrationkit.GooseMigrationOpt{
		Dir:      "/database/migrations",
		DBString: mysqlkit.GetConnUrl(&env.MysqlConnConfig),
	})
	if err := migrationKit.Up(); err != nil {
		logkit.Fatal(ctx, "db migration failed", logkit.Payload{"err": err})
	}
	migrationKit.Close()

	// init validator
	logkit.Infof(ctx, "init validator")
	validator := validatorkit.NewGoPlaygroundValidator()

	// init server base
	logkit.Infof(ctx, "init server")
	serv := rpc.NewBevisChangServer(rpc.BevisChangServerOpt{
		Validator: validator,
		RecordDao: dao.NewRecordDAO(db, cacheSrv),
	})

	// init service
	var wg sync.WaitGroup

	launchers := []*ServiceLauncher{
		NewGrpcSvcLauncher(env.GRPCAddr, serv),
		NewHttpSvcLauncher(env.HTTPAddr, serv),
	}

	logkit.Infof(ctx, "launching service")

	// launch service
	for i := range launchers {
		l := launchers[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := l.Run(); err != nil {
				logkit.Fatal(ctx, "launcher failed", logkit.Payload{"err": err, "labels": l.Labels})
			}
		}()
	}

	wg.Wait()
}

// NewGrpcSvcLauncher 3-1. You need add a gRPC listener and register the service.
func NewGrpcSvcLauncher(addr string, serv pb.BevisChangServer) *ServiceLauncher {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logkit.Fatal(context.TODO(), "failed to start listen tpc", logkit.Payload{"err": err})
	}

	s := grpc.NewServer()

	pb.RegisterBevisChangGrpcService(s, serv) // 3-2. Run "RegisterBevisChangGrpcService"

	return &ServiceLauncher{
		Labels: []string{"grpc"},
		Run: func() error {
			return s.Serve(lis)
		},
	}
}

// NewHttpSvcLauncher 4-1. You need add a HTTP listener and register the service.
func NewHttpSvcLauncher(addr string, serv pb.BevisChangServer) *ServiceLauncher {

	// TODO: move details into RegisterBevisChangHttpService

	s := gin.New()
	s.Use(gin.Recovery())
	s.Use(metrickit.Middleware(metrickit.New("gin")))

	pb.RegisterBevisChangHttpService(s, serv) // 4-2. Run "RegisterBevisChangHttpService"

	return &ServiceLauncher{
		Labels: []string{"http"},
		Run: func() error {
			return s.Run(addr)
		},
	}
}
