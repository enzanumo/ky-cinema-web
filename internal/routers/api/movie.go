package api

import (
	"github.com/enzanumo/ky-theater-web/internal/core/service"
	"github.com/enzanumo/ky-theater-web/pkg/app"
	"github.com/enzanumo/ky-theater-web/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetMovieList(c *gin.Context) {
	response := app.NewResponse(c)

	ds := service.DS

	li, err := ds.GetMovieList()

	if err != nil {
		logrus.Errorf("service.GetMovieList err: %v\n", err)
		response.ToErrorResponse(errcode.GetCollectionsFailed)
		return
	}

	response.ToResponseList(li, int64(len(li)))
}
