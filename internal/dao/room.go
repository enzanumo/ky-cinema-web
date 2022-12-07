package dao

import "github.com/enzanumo/ky-theater-web/internal/model"

func (s *dataServant) GetRoomList() (rooms []*model.Room, err error) {
	tx := s.db.Find(&rooms)
	err = tx.Error
	return
}
