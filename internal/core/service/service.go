package service

import (
	"github.com/enzanumo/ky-theater-web/internal/core"
	"github.com/enzanumo/ky-theater-web/internal/dao"
)

var (
	DS core.DataService
)

func Initialize() {
	DS = dao.DataService()
}
