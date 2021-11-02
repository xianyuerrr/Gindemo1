package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	customRoute(r)
	if err := r.Run(":8000"); err != nil {
		return
	}
}
