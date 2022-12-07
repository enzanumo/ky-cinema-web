package model

import "time"

//Schedule 场次
type Schedule struct {
	StartAt time.Time `gorm:"uniqueIndex:schedule_unique"`
	RoomID  int       `gorm:"uniqueIndex:schedule_unique"`
}

//Seat 一个场次中的一个座位 票的最小单位
type Seat struct {
	Schedule
	SeatIdx int64
}

//Ticket 关系，这个座位卖给了谁
type Ticket struct {
	Seat
	UserID      int64
	OrigPrice   Price
	ActualPrice Price
}

type TicketFormatted struct {
	Ticket
	Movie *Movie
	EndAt time.Time
}

//TicketPlan 一次上映
type TicketPlan struct {
	Schedule
	MovieID   int
	EndAt     time.Time
	BasePrice Price
}
