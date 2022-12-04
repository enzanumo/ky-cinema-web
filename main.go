package main

import (
	"fmt"
	"github.com/enzanumo/ky-theater-system/internal/routers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	gin.SetMode("debug")

	addr := "0.0.0.0:9091"

	router := routers.NewRouter()
	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("ky-theater-system service listen on http://%s\n", addr)
	_ = s.ListenAndServe()
}
