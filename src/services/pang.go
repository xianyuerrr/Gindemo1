package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pang(r *gin.Context) {
	r.JSON(http.StatusOK, gin.H{"message": "pang", "haha": "haha"})
}
