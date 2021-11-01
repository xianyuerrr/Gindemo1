package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkConfig(cq map[string][]string) bool {
	//todo
	return true
}
func Config(c *gin.Context) {
	platform := c.PostForm("platform")                     //string
	downloadUrl := c.PostForm("download_url")              //string
	updateVersionCode := c.PostForm("update_version_code") //string
	md5 := c.PostForm("md5")                               //string
	//deviceIdList := c.Query("device_id_list")                  //string
	//maxUpdateVersionCode := c.Query("max_update_version_code") //string
	//minUpdateVersionCode := c.Query("min_update_version_code") //string
	//maxOsApi := c.Query("max_os_api")                          //int
	//minOsApi := c.Query("min_os_api")                          //int
	//cpuArch := c.Query("cpu_arch")                             //string
	//channel := c.Query("channel")                              //string
	//title := c.Query("title")                                  //string
	//updateTips := c.Query("update_tips")                       //string
	fmt.Println(platform, downloadUrl, updateVersionCode, md5)
	if checkConfig(c.Request.URL.Query()) {
		c.String(http.StatusOK, "post success")
	} else {
		c.String(http.StatusBadGateway, "check 	fail")
	}

}
