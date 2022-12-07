package internal

import (
	"github.com/enzanumo/ky-theater-web/internal/model/migration"
	"github.com/enzanumo/ky-theater-web/internal/routers/api"
	"github.com/enzanumo/ky-theater-web/internal/service"
)

func Initialize() {
	// migrate database if needed
	migration.Run()

	// initialize service
	service.Initialize()
	api.Initialize()
}
