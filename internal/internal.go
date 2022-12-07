package internal

import (
	"github.com/enzanumo/ky-theater-web/internal/core/service"
	"github.com/enzanumo/ky-theater-web/internal/routers/api"
)

func Initialize() {
	// initialize service
	service.Initialize()
	api.Initialize()
}
