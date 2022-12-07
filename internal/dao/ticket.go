package dao

import (
	"fmt"
	"github.com/enzanumo/ky-theater-web/internal/model"
	"gorm.io/gorm"
	"time"
)

var forever time.Time = time.Unix(1<<62, 0)

func (s *dataServant) GetAvailableSeats(schedule *model.Schedule) ([]int8, error) {
	room := model.Room{ID: schedule.RoomID}
	err := s.db.Take(&room).Error
	if err != nil {
		return nil, err
	}

	res := make([]int8, room.SeatCount)

	rows, err := s.db.Model(&model.Ticket{}).Where(
		"start_at = ? and room_id = ? and expire_at > CURRENT_TIMESTAMP",
		schedule.StartAt, schedule.RoomID,
	).Rows()
	defer rows.Close()

	for rows.Next() {
		var seatIdx int64
		s.db.ScanRows(rows, &seatIdx)
		res[seatIdx] = 1
	}

	return res, nil
}

func (s *dataServant) LockSeat(seat *model.Seat, user *model.User) error {
	tck := model.Ticket{}
	err := s.db.Take(&tck, seat).Error

	ok := false

	if err == gorm.ErrRecordNotFound {
		ok = true
	} else {
		return err
	}
	if time.Now().After(tck.ExpireAt) {
		ok = true
	}

	if ok {
		tck.Seat = *seat
		tck.ExpireAt = time.Now().Add(10 * time.Minute)
		s.db.Create(&tck)
		return nil
	}
	return fmt.Errorf("lock failed")
}

func (s *dataServant) GrantSeat(seat *model.Seat) error {
	tck := model.Ticket{
		Seat: *seat,
	}
	err := s.db.Model(&tck).Update("expire_at", forever).Error
	return err
}

func (s *dataServant) SetSeatPromotion(schedule *model.Schedule, seatIdxes []int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *dataServant) AddSchedule(sp *model.TicketPlan) error {
	err := s.db.Create(sp).Error
	return err
}

func (s *dataServant) GetUserTickets(user *model.User) (tcks []*model.Ticket, err error) {
	err = s.db.Where("user_id = ?", user.ID).Find(&tcks).Error
	return
}
