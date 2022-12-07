package service

import (
	"github.com/enzanumo/ky-theater-web/internal/core"
	"github.com/enzanumo/ky-theater-web/internal/dao"
)

var (
	ds core.DataService
)

func Initialize() {
	ds = dao.DataService()
}
