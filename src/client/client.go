package client

import (
	"fmt"
	"grayRelease/src/router/middleware/authencator"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func generateUrl(action string) string {
	user := "admin"
	password := "admin"
	baseURL := "127.0.0.1:8000/config/" + action
	createTime := time.Now().Unix()
	authToken := authencator.CreateAuthToken(baseURL, user, password, createTime)
	token := authToken.GetToken()
	return "http://" + fmt.Sprintf("%s?user=%s&token=%s&create_time=%d", baseURL, user, token, createTime)
}

func main() {
	action := [...]string{
		// "release",
		"offline",
	}
	formData := [...]url.Values{
		url.Values{"id": {"2"}},
		url.Values{"id": {"2"}},
	}

	for i := 0; i < len(action); i++ {
		reqUrl := generateUrl(action[i])
		resp, err := http.PostForm(reqUrl, formData[i])
		if err != nil {
			log.Fatal(err)
		}
		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", bodyText)
	}
}
