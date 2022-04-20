package main

import (
	"github.com/AmazingTalker/go-rpc-kit/flagkit"
	"github.com/AmazingTalker/go-rpc-kit/logkit"
	"github.com/AmazingTalker/go-rpc-kit/mysqlkit"
)

type RedisConfig struct {
	// Addrs indicates the map of name => host:port addresses of ring shards.
	Addrs map[string]string `long:"addrs" description:"a map from name to the pair of host and port" env:"ADDRS" env-delim:","`
}

// LocalCacheConfig means the config for local cache
type LocalCacheConfig struct {
	// default size is 67108864 = 64*1024*1024
	Size int `long:"size" description:"the size of the cache" default:"67108864" env:"SIZE"`
}

type AirbrakeConfig struct {
	ProjectID  int64  `long:"projectId" default:"" env:"PROJECT_ID"`
	ProjectKey string `long:"projectKey" default:"" env:"PROJECT_KEY"`
}

type LoggerConfig struct {
	Level       logkit.LoggerLevel `long:"level" description:"set log level" default:"info" env:"LEVEL"`
	Development bool               `long:"development" description:"enable development mode" env:"DEVELOPMENT"`
	Airbrake    AirbrakeConfig     `group:"airbrake" namespace:"airbrake" env-namespace:"AIRBRAKE"`
}

type EnvConfig struct {
	Namespace   string `long:"namespace" description:"Environment namespace. ex: Prod, Stag, Dev" env:"NAMESPACE"`
	PodName     string `long:"pod" description:"pod name or host name in k8s" env:"POD_NAME"`
	ServiceName string `long:"service" description:"service name" env:"SERVICE_NAME"`
}

type MetricConfig struct {
	URL            string  `long:"url" description:"metric URL" env:"URL"`
	APIKey         string  `long:"apikey" description:"api key" env:"API_KEY"`
	RefleshSeconds float64 `long:"refleshseconds" description:"reflesh seconds" env:"REFLESH_SECONDS"`
}

type MonitorConfig struct {
	PeriodSecs uint `long:"period-in-secs" description:"period of collecting in seconds" default:"10" env:"PERIOD_SECONDS"`
}

type EtcdConfig struct {
	Addrs              []string `long:"addrs" description:"A slice of strings includs the pair of host and port" env:"ADDRS" env-delim:","`
	DialTimeoutSeconds int      `long:"dialTimeoutSeconds" description:"timeout seconds for dialing at the beginning" env:"DIAL_TIMEOUT_SECONDS"`
}

var env struct {
	HTTPAddr                 string `short:"h" long:"http.addr" env:"HTTP_ADDR" default:":8080"`
	GRPCAddr                 string `short:"g" long:"grpc.addr" env:"GRPC_ADDR" default:":8081"`
	LoggerConfig             `group:"logger" namespace:"logger" env-namespace:"LOGGER"`
	mysqlkit.MysqlConnConfig `group:"mysql" namespace:"mysql" env-namespace:"MYSQL"`
	RedisConfig              `group:"redis" namespace:"redis" env-namespace:"REDIS"`
	LocalCacheConfig         `group:"lc" namespace:"lc" env-namespace:"LOCAL_CACHE"`
	EnvConfig                `group:"env" namespace:"env" env-namespace:"ENV"`
	MetricConfig             `group:"metric" namespace:"metric" env-namespace:"METRIC"`
	MonitorConfig            `group:"monitor" namespace:"monitor" env-namespace:"MONITOR"`
	EtcdConfig               `group:"etcd" namespace:"etcd" env-namespace:"ETCD"`
}

func init() {
	flagkit.Register(&env)
}
