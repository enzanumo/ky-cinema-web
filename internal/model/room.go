package model

import "gorm.io/gorm"

type Room struct {
	ID        int64 `gorm:"primarykey"`
	Name      string
	SeatCount int
	Tags      string
}

func (m *Room) Create(db *gorm.DB) (*Room, error) {
	err := db.Create(&m).Error
	return m, err
}
