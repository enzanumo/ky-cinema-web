package core

import "github.com/enzanumo/ky-theater-web/internal/model"

type RoomService interface {
	GetRoomList() (rooms []*model.Room, err error)
}
