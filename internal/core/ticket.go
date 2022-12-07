package core

import "github.com/enzanumo/ky-theater-web/internal/model"

type TicketService interface {
	GetAvailableSeats(schedule *model.Schedule) ([]int8, error)
	LockSeat(seat *model.Seat, user *model.User) error
	GrantSeat(seat *model.Seat) error
	SetSeatPromotion(schedule *model.Schedule, seatIdxes []int64) error

	AddSchedule(sp *model.TicketPlan) error
	GetMovieSchedule(movie *model.Movie) ([]*model.TicketPlan, error)
	GetUserTickets(user *model.User) ([]*model.Ticket, error)
}
