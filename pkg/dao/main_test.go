package dao

import (
	"context"
	"os"
	"testing"

	"github.com/AmazingTalker/go-rpc-kit/logkit"
	"github.com/go-redis/redis/v8"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"

	"github.com/AmazingTalker/go-cache"
	"github.com/AmazingTalker/go-rpc-kit/cachekit"
	"github.com/AmazingTalker/go-rpc-kit/migrationkit"
	"github.com/AmazingTalker/go-rpc-kit/mysqlkit"
	"github.com/AmazingTalker/go-rpc-kit/rediskit"
)

func getMigrationDir() string {
	dir := os.Getenv("MIGRATION_DIR")
	if dir != "" {
		return dir
	}
	return "/src/database/migrations"
}

func getConnUrl() string {
	envUrl := os.Getenv("MYSQL_DSN")
	if envUrl != "" {
		return envUrl
	}
	return "root:root@tcp(localhost:3306)/go-amazing"
}

func getRingAddrs() map[string]string {
	return map[string]string{"server1": ":6379"}
}

func setupDB() {
	migrationKit := migrationkit.NewGooseMigrationKit(migrationkit.GooseMysqlDriver, migrationkit.GooseMigrationOpt{
		Dir:      getMigrationDir(),
		DBString: getConnUrl(),
	})
	if err := migrationKit.Up(); err != nil {
		panic(err)
	}
	migrationKit.Close()
}

var testCtx context.Context
var testConn *gorm.DB
var testRing *redis.Ring
var testCache cache.Service

var _ = Describe("MysqlDAO", func() {
	BeforeSuite(func() {
		setupDB()

		logkit.RegisterAmazingLogger(&logkit.Config{
			Logger:      logkit.LoggerZap,
			Level:       logkit.PanicLevel,
			Development: false, // turn it on if you want more gorm information when running tests.
			IntegrationAirbrake: &logkit.IntegrationAirbrake{
				ProjectID:  0,
				ProjectKey: "",
				Env:        "ci",
			},
		})

		testCtx = context.Background()
		testConn, _ = mysqlkit.NewMySqlConn(mysqlkit.MySqlConnOpt{
			Config: &mysqlkit.MysqlConnConfig{
				DSN: getConnUrl(),
			},
		})

		testRing, _ = rediskit.NewRedisRing(getRingAddrs())
		_ = testRing.ForEachShard(testCtx, func(ctx context.Context, client *redis.Client) error {
			return client.FlushDB(ctx).Err()
		})

		testCache = cachekit.NewCache(cachekit.NewSharedCache(testRing), cachekit.NewLocalCache(1024))
	})
})

func TestPGRecordDAO(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DAO Suite")
}
