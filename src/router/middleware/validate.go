package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// session/cookie中获取
		username := c.Query("user")
		password := c.Query("passwd")

		if username == "admin" && password == "admin" {
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "身份验证失败"})
			return
		}
	}
}
