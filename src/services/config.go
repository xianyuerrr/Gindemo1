package services

import (
	"demo1/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkConfig(config model.Rule) bool {
	//todo

	if model.AddRule(config) == 0 {
		return true
	}
	return false

}
func Config(c *gin.Context) {

	var config model.Rule
	if err := c.Bind(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if checkConfig(config) {
		c.String(http.StatusOK, "post success")
	} else {
		c.String(http.StatusBadRequest, "check fail")
	}
}
