package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	customRoute(r)
	r.Run(":8000")
}
