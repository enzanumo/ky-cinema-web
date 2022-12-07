package routers

import (
	"github.com/enzanumo/ky-theater-web/internal/middleware"
	"github.com/enzanumo/ky-theater-web/internal/routers/api"
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

	// 注册 静态资源 路由
	{
		registerStatick(e)
	}

	// v1 group api
	r := e.Group("/v1")

	// 获取version
	r.GET("/", api.Version)

	// 用户登录
	r.POST("/auth/login", api.Login)

	// 用户注册
	r.POST("/auth/register", api.Register)

	// 无鉴权路由组
	noAuthApi := r.Group("/")
	{
		// 获取用户基本信息
		noAuthApi.GET("/user/profile", api.GetUserProfile)

		// 获取电影信息
		r.GET("/movies", api.Stub)

		// 获取电影详情
		r.GET("/movies/:id", api.Stub)

		// 获取电影排期
		r.GET("/movies/:id/schedule", api.Stub)

		// 查询场次余票
		r.POST("/seats", api.Stub)
	}

	// 鉴权路由组
	authApi := r.Group("/").Use(middleware.JWT())
	adminApi := r.Group("/").Use(middleware.JWT()).Use(middleware.Admin())
	{
		// 获取当前用户信息
		authApi.GET("/user/info", api.GetUserInfo)
		// 绑定用户手机号
		{
			authApi.POST("/user/phone", api.BindUserPhone)
		}
		// 修改密码
		authApi.POST("/user/password", api.ChangeUserPassword)
		// 修改昵称
		authApi.POST("/user/nickname", api.ChangeNickname)
		// 下单
		authApi.POST("/orders/create", api.Stub)
		// 支付
		authApi.POST("/orders/pay", api.Stub)
		// 获取我的订单
		authApi.GET("/orders", api.Stub)
		// 获取我的订单详情
		authApi.GET("/orders/details/:id", api.Stub)
		// 获取我的信息
		authApi.GET("/user/info")

		// 管理
		// 新增影片
		adminApi.POST("/movies", api.Stub)

		// 获取所有场次
		adminApi.GET("/schedule", api.Stub)

		// 新增场次
		adminApi.POST("/schedule", api.Stub)

		// 新增场次优惠
		adminApi.POST("/schedule/promotion", api.Stub)

		// 查询用户信息
		adminApi.GET("/user", api.Stub)
		// 添加用户
		adminApi.POST("/user", api.Stub)
		// 修改用户
		adminApi.POST("/user/:id", api.Stub)
	}

	// 默认404
	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Not Found",
		})
	})

	// 默认405
	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  "Method Not Allowed",
		})
	})

	return e
}
