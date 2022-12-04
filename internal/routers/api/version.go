package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, "1.0")
}
func Stub(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "NotImplemented")
}
