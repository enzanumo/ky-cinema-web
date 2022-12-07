package model

import "time"

type Schedule struct {
	StartAt time.Time `gorm:"uniqueIndex:schedule_unique"`
	RoomID  int       `gorm:"uniqueIndex:schedule_unique"`
}

type PlannedSchedule struct {
	Schedule
	MovieID   int
	EndAt     time.Time
	BasePrice Price
}
