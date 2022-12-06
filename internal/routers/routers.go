package routers

import (
	"github.com/enzanumo/ky-theater-system/internal/routers/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	e := gin.New()
	e.HandleMethodNotAllowed = true
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	// 跨域配置
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	e.Use(cors.New(corsConfig))

	// v1 group api
	r := e.Group("/v1")

	// 获取version
	r.GET("/", api.Version)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	r.GET("/movies", api.Movies)

	r.GET("/movies/:id", api.Stub)
	r.GET("/movies/:id/schedule", api.Stub)

	r.POST("/seats", api.Stub)

	self := r.Group("/")

	self.POST("/orders/create", api.Stub)
	self.POST("/orders/pay", api.Stub)
	self.GET("/orders", api.Stub)
	self.GET("/orders/details/:id", api.Stub)

	self.GET("/user/info", api.Stub)

	adm := r.Group("/")

	adm.POST("/movies", api.Stub)
	adm.GET("/schedule", api.Stub)
	adm.POST("/schedule", api.Stub)
	adm.POST("/schedule/promotion", api.Stub)

	adm.GET("/user", api.Stub)
	adm.POST("/user", api.Stub)
	adm.POST("/user/:id", api.Stub)

	return e
}
