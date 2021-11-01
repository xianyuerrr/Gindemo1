package services

import "github.com/gin-gonic/gin"

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
