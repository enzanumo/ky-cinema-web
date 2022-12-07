package conf

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

type Setting struct {
	vp *viper.Viper
}

type LoggerSettingS struct {
	Level string
}

type LoggerFileSettingS struct {
	SavePath string
	FileName string
	FileExt  string
}

type ServerSettingS struct {
	RunMode      string
	HttpIp       string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DatabaseSetingS struct {
	Use         string
	TablePrefix string
	LogLevel    string
}

type MySQLSettingS struct {
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type Sqlite3SettingS struct {
	Path string
}

type RedisSettingS struct {
	Host     string
	Password string
	DB       int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath(".")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}

func (s *Setting) Unmarshal(objects map[string]interface{}) error {
	for k, v := range objects {
		err := s.vp.UnmarshalKey(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *MySQLSettingS) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		s.UserName,
		s.Password,
		s.Host,
		s.DBName,
		s.Charset,
		s.ParseTime,
	)
}

func (s *Sqlite3SettingS) Dsn(driverName string) string {
	pragmas := "_foreign_keys=1&_journal_mode=WAL&_synchronous=NORMAL&_busy_timeout=8000"
	if driverName == "sqlite" {
		pragmas = "_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(8000)&_pragma=journal_size_limit(100000000)"
	}
	return fmt.Sprintf("file:%s?%s", s.Path, pragmas)
}

func (s *DatabaseSetingS) logLevel() logger.LogLevel {
	switch strings.ToLower(s.LogLevel) {
	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Error
	}
}

func (s *LoggerSettingS) logLevel() logrus.Level {
	switch strings.ToLower(s.Level) {
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.ErrorLevel
	}
}
func endpoint(host string, secure bool) string {
	schema := "http"
	if secure {
		schema = "https"
	}
	return schema + "://" + host
}
