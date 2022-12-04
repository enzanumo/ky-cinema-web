package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Movies(ctx *gin.Context) {
	mvs := []gin.H{
		{
			"name": "aaa",
		},
	}
	ctx.JSON(http.StatusOK, RspOK(gin.H{
		"movies": mvs,
	}))
}
