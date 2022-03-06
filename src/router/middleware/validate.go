package middleware

import (
	"github.com/gin-gonic/gin"
	"grayRelease/src/conf"
	"grayRelease/src/router/middleware/authencator"
	"net/http"
)

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var apiRequest *authencator.ApiRequest
		var credentialStorage authencator.CredentialStorage
		var apiAuthenticator authencator.ApiAuthenticator

		authorConfig := conf.GetMysqlAuthorConfig()
		apiRequest = authencator.CreatFromFullUrl(c.Request.Host + c.Request.RequestURI)
		credentialStorage = authencator.GetMysqlCredentialStorage(authorConfig.Dsn, authorConfig.MaxIdle, authorConfig.MaxOpen)
		apiAuthenticator = authencator.GeDefaulttApiAuthencator(credentialStorage)
		if apiRequest != nil && apiAuthenticator.Auth(*apiRequest) {
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": "身份验证失败"})
			return
		}
	}
}
