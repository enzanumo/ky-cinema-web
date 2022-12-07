package model

import "time"

//Schedule 场次
type Schedule struct {
	StartAt *time.Time
	RoomID  int64
}

//Seat 一个场次中的一个座位 票的最小单位
type Seat struct {
	StartAt *time.Time `gorm:"uniqueIndex:seat_unique"`
	RoomID  int64      `gorm:"uniqueIndex:seat_unique"`
	SeatIdx int64
}

//Ticket 关系，这个座位卖给了谁
type Ticket struct {
	Seat
	UserID   int64
	ExpireAt *time.Time
}

type TicketFormatted struct {
	Ticket
	Movie *Movie
	EndAt *time.Time
}

//TicketPlan 一次上映
type TicketPlan struct {
	StartAt   *time.Time `gorm:"uniqueIndex:ticket_plan_unique"`
	RoomID    int64      `gorm:"uniqueIndex:ticket_plan_unique"`
	MovieID   int
	EndAt     *time.Time
	BasePrice Price
}
