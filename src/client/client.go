package client

import (
	"fmt"
	"grayRelease/src/client/operationTest"
	"grayRelease/src/router/middleware/authencator"
	"log"
	"net/http"
	"time"
)

func GenerateUrl(action string, needAuth bool) string {
	if !needAuth {
		return "http://127.0.0.1:8000/check"
	}
	baseURL := "127.0.0.1:8000/config/" + action
	user := "admin"
	password := "admin"
	createTime := time.Now().Unix()
	authToken := authencator.CreateAuthToken(baseURL, user, password, createTime)
	token := authToken.GetToken()
	return "http://" + fmt.Sprintf("%s?user=%s&token=%s&create_time=%d", baseURL, user, token, createTime)
}

func testOperation(test operationTest.Operation) *http.Response {
	url := GenerateUrl(test.Action, test.NeedAuth)
	resp, err := http.PostForm(url, test.FormData)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return resp
}
