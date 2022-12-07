package api

import (
	"github.com/enzanumo/ky-theater-web/pkg/app"
	"github.com/gin-gonic/gin"
)

func Initialize() {
}

func Version(c *gin.Context) {
	response := app.NewResponse(c)
	response.ToResponse(gin.H{
		"BuildInfo": "1.0",
	})
}
