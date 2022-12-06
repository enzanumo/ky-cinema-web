package conf

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBEngine *gorm.DB
	Redis    *redis.Client
)

func SetupDBEngine() (err error) {

	DBEngine, err = gorm.Open(sqlite.Open("ky-cinema.sqlite3"))

	if err != nil {
		return err
	}

	Redis = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1",
	})

	return nil
}
