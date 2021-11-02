package main

import (
	"demo1/src/controller"
	"github.com/gin-gonic/gin"
)

func customRoute(r *gin.Engine) {
	r.GET("/check", controller.Check)
	r.POST("/config", controller.Config)
}
