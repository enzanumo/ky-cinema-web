package conf

import (
	"github.com/enzanumo/ky-theater-web/internal/model/migrate"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var (
	db    *gorm.DB
	Redis *redis.Client
	once  sync.Once
)

func MustGormDB() *gorm.DB {
	once.Do(func() {
		var err error
		if db, err = newDBEngine(); err != nil {
			panic(err)
			//logrus.Fatalf("new gorm db failed: %s", err)
		}
	})
	return db
}

func newDBEngine() (*gorm.DB, error) {
	newLogger := logger.New(
		logrus.StandardLogger(), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second,                // 慢 SQL 阈值
			LogLevel:                  DatabaseSetting.logLevel(), // 日志级别
			IgnoreRecordNotFoundError: true,                       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,                      // 禁用彩色打印
		},
	)

	config := &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   DatabaseSetting.TablePrefix,
			SingularTable: true,
		},
	}

	plugin := dbresolver.Register(dbresolver.Config{}).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(MysqlSetting.MaxIdleConns).
		SetMaxOpenConns(MysqlSetting.MaxOpenConns)

	var (
		db  *gorm.DB
		err error
	)
	if useMysql() {
		logrus.Debugln("use MySQL as db")
		if db, err = gorm.Open(mysql.Open(MysqlSetting.Dsn()), config); err == nil {
			db.Use(plugin)
		}
	} else {
		logrus.Debugf("use Sqlite3 as db path:%s", Sqlite3Setting.Path)
		db, err = gormOpenSqlite3(config)
	}

	if useMigrate() {
		migrate.Run(db)
	}

	return db, err
}

func useMigrate() bool {
	return true
}

func useMysql() bool {
	return strings.ToLower(DatabaseSetting.Use) == "mysql"
}

func setupDBEngine() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     redisSetting.Host,
		Password: redisSetting.Password,
		DB:       redisSetting.DB,
	})
}
