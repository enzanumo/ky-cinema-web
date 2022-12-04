package api

import "github.com/gin-gonic/gin"

//func Initalize() {
//
//}

func Rsp(code int, msg string, h gin.H) gin.H {
	return gin.H{
		"code":  code,
		"msg":   msg,
		"value": h,
	}
}

func RspOK(h gin.H) gin.H {
	return gin.H{
		"code":  0,
		"msg":   "ok",
		"value": h,
	}
}
