package migrate

import (
	"github.com/enzanumo/ky-theater-web/internal/model"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(
		&model.Movie{},
		&model.Order{},
		&model.Room{},
		&model.Schedule{},
		&model.User{},
	)
}
