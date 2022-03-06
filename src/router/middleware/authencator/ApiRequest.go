package authencator

import (
	"fmt"
	"strconv"
	"strings"
)

type ApiRequest struct {
	baseURL    string
	token      string
	user       string
	createTime int64 // 1504079553
}

func (apiRequest ApiRequest) getBaseUrl() string {
	return apiRequest.baseURL
}

func (apiRequest ApiRequest) getToken() string {
	return apiRequest.token
}

func (apiRequest ApiRequest) getAppId() string {
	return apiRequest.user
}

func (apiRequest ApiRequest) getTimeStamp() int64 {
	return apiRequest.createTime
}

func CreatFromFullUrl(url string) *ApiRequest {
	if url == "" {
		return nil
	}
	idx := strings.Index(url, "?")
	if idx == -1 {
		fmt.Println("failure to create ApiRequest")
		return nil
	}
	baseUrl := url[:idx]
	paramsStr := url[idx+1:]
	params := make(map[string]string)
	paramsList := strings.Split(paramsStr, "&")
	for i := 0; i < len(paramsList); i++ {
		split := strings.Split(paramsList[i], "=")
		if len(split) == 1 {
			fmt.Println("failure to create ApiRequest")
			return nil
		}
		key := split[0]
		value := split[1]
		params[key] = value
	}

	token := params["token"]
	user := params["user"]
	timestampStr := params["create_time"]
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return nil
	}
	return &ApiRequest{baseUrl, token, user, timestamp}
}

// 获取时间戳
// timestamp := time.Now().Unix() // 1504079553

// 时间戳转Time 再转 string
// timeNow := time.Unix(timestamp, 0) //2017-08-30 16:19:19 +0800 CST
// timeString := timeNow.Format("2006-01-02 15:04:05") //2015-06-15 08:52:32

// string 转 时间戳
// stringTime := "2017-08-30 16:40:41"
// loc, _ := time.LoadLocation("Local")
// the_time, err := time.ParseInLocation("2006-01-02 15:04:05", stringTime, loc)
// if err == nil {
// unix_time := the_time.Unix() // 1504082441
// }
