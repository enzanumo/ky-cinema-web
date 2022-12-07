package migrate

import (
	"github.com/enzanumo/ky-theater-web/internal/model"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	db.AutoMigrate(&model.Movie{})
	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.Room{})
	db.AutoMigrate(&model.Ticket{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.TicketPlan{})
	panic("done")
}
