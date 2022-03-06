package authencator

import (
	"fmt"
	"testing"
	"time"
)

func TestAuthToken_GetToken(t *testing.T) {
	// url := "127.0.0.1:8000?user=admin&token=1&create_time="
	baseURL := "127.0.0.1:8000/config/release"
	user := "admin"
	password := "admin"
	createTime := time.Now().Unix()
	authToken1 := CreateAuthToken(baseURL, user, password, createTime)
	authToken2 := CreateAuthToken(baseURL, user, password, createTime)
	fmt.Println(authToken1)
	fmt.Println(authToken2)
}
