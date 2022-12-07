package core

import "github.com/enzanumo/ky-theater-web/internal/model"

type TicketService interface {
	GetAvailableSeats(schedule *model.Schedule) ([]int8, error)
	LockSeat(schedule *model.Schedule, user *model.User, seatIdx int64) error
	GrantSeat(schedule *model.Schedule, user *model.User, seatIdx int64) error
	SetSeatPromotion(schedule *model.Schedule, seatIdxes []int64) error

	AddSchedule(sp *model.TicketPlan) error

	GetUserTickets(user *model.User) ([]*model.Ticket, error)
}
