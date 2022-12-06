package model

type Order struct {
	*Model
	Sched       Schedule `gorm:"embedded;embeddedPrefix:schedule_"`
	SeatIndex   int
	OrigPrice   Price
	ActualPrice Price
}
