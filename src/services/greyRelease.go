package services

import (
	"demo1/src/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

func Check(c *gin.Context) {
	var respUrl string
	version := c.Query("version")
	device_platform := c.Query("device_platform")
	device_id := c.Query("device_id")
	os_api := c.Query("os_api")
	channel := c.Query("channel")
	version_code := c.Query("version_code")
	update_version_code := c.Query("update_version_code")
	aid := c.Query("aid")
	cpu_arch := c.Query("cpu_arch")
	rules := model.GetRules()

	for index := 0; index < len(*rules); index++ {
		//if rule continue
		respUrl = (*rules)[index].GreyLink
		break
	}
	c.JSON(http.StatusOK, gin.H{"url:": respUrl})
}
func Config(c *gin.Context) {
	platform := c.Query("platform")                            //string
	downloadUrl := c.Query("download_url")                     //string
	updateVersionCode := c.Query("update_version_code")        //string
	md5 := c.Query("md5")                                      //string
	deviceIdList := c.Query("device_id_list")                  //string
	maxUpdateVersionCode := c.Query("max_update_version_code") //string
	minUpdateVersionCode := c.Query("min_update_version_code") //string
	maxOsApi := c.Query("max_os_api")                          //int
	minOsApi := c.Query("min_os_api")                          //int
	cpuArch := c.Query("cpu_arch")                             //string
	channel := c.Query("channel")                              //string
	title := c.Query("title")                                  //string
	updateTips := c.Query("update_tips")                       //string

}
