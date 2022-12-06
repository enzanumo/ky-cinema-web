package model

import "time"

type Schedule struct {
	StartAt time.Time `gorm:"uniqueIndex:schedule_unique"`
	Room    *Room     `gorm:"uniqueIndex:schedule_unique"`
}

type PlannedSchedule struct {
	*Schedule
	Movie     *Movie
	EndAt     time.Time
	BasePrice Price
}
