package api

import (
	"github.com/enzanumo/ky-theater-web/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Initialize() {
}

func Version(c *gin.Context) {
	response := app.NewResponse(c)
	response.ToResponse(gin.H{
		"BuildInfo": "1.0",
	})
}

func Stub(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "NotImplemented")
}
