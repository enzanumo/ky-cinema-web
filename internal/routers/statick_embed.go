package routers

import (
	"net/http"

	"github.com/enzanumo/ky-theater-web/web"
	"github.com/gin-gonic/gin"
)

// registerStatick register static assets route
func registerStatick(e *gin.Engine) {
	routeStatic(e, "/", "/index.html", "/favicon.ico", "/assets/*filepath")
}

func routeStatic(e *gin.Engine, paths ...string) {
	staticHandler := http.FileServer(web.NewFileSystem())
	handler := func(c *gin.Context) {
		staticHandler.ServeHTTP(c.Writer, c.Request)
	}
	for _, path := range paths {
		e.GET(path, handler)
		e.HEAD(path, handler)
	}
}
