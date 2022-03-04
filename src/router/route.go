package router

import (
	"github.com/gin-gonic/gin"
	"grayRelease/src/controller"
	"grayRelease/src/router/middleware"
)

func CustomRoute(r *gin.Engine) {
	r.POST("/check", controller.Check)
	r.POST("/config/:action",
		middleware.Validate(),
		func(c *gin.Context) { controller.Config(c) },
	)
}
