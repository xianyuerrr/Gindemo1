package router

import (
	"github.com/gin-gonic/gin"
	"grayRelease/src/controller"
	"grayRelease/src/router/middleware"
	"net/http"
)

func CustomRoute(r *gin.Engine) {
	r.POST("/check", controller.Check)
	r.POST("/config/:action",
		middleware.Validate(),
		func(c *gin.Context) { ConfigRouter(c) },
	)
}

func ConfigRouter(c *gin.Context) {
	action := c.Param("action")
	// db := model.GetWriteDb()
	// db.Begin()
	var res bool
	if action == "" {
		// writes the given string into the response body.
		c.String(http.StatusBadRequest, "There is not a action named nil.")
		return
	} else if action == "add" {
		res = controller.AddRule(c)
	} else if action == "delete" {
		res = controller.DeleteRule(c)
	} else if action == "update" {
		res = controller.UpdateRule(c)
	} else if action == "release" {
		res = controller.ReleaseRule(c)
	} else if action == "offline" {
		res = controller.OfflineRule(c)
	}
	if res {
		// db.Commit()
	} else {
		// db.Rollback()
	}
}
