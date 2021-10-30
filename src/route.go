package main

import (
	"demo1/src/services"
	"github.com/gin-gonic/gin"
)

func customRoute(r *gin.Engine) {
	r.GET("/ping", services.Pang)
	r.GET("/", services.Hit)
}
