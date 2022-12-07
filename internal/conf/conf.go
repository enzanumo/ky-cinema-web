package conf

import (
	"log"
	"time"
)

var (
	loggerSetting     *LoggerSettingS
	loggerFileSetting *LoggerFileSettingS
	redisSetting      *RedisSettingS

	DatabaseSetting *DatabaseSetingS
	MysqlSetting    *MySQLSettingS
	Sqlite3Setting  *Sqlite3SettingS
	ServerSetting   *ServerSettingS
	JWTSetting      *JWTSettingS
)

func setupSetting(suite []string, noDefault bool) error {
	setting, err := NewSetting()
	if err != nil {
		return err
	}

	objects := map[string]interface{}{
		"Server":     &ServerSetting,
		"Logger":     &loggerSetting,
		"LoggerFile": &loggerFileSetting,
		"Database":   &DatabaseSetting,
		"MySQL":      &MysqlSetting,
		"Sqlite3":    &Sqlite3Setting,
		"Redis":      &redisSetting,
		"JWT":        &JWTSetting,
	}
	if err = setting.Unmarshal(objects); err != nil {
		return err
	}

	JWTSetting.Expire *= time.Second
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second

	return nil
}

func Initialize(suite []string, noDefault bool) {
	err := setupSetting(suite, noDefault)
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	setupLogger()
	setupDBEngine()
}
