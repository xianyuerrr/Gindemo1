package authencator

import (
	"fmt"
	"testing"
)

var dsn = "root:123456@tcp(localhost:3306)/author?charset=utf8mb4&parseTime=True&loc=Local"
var maxIdle = 100
var maxOpen = 10

// var lifeTime = "+3s"
// var log = true

func TestMysqlCredentialStorage_AddAuthor(t *testing.T) {
	s := GetMysqlCredentialStorage(dsn, maxIdle, maxOpen)
	res := s.AddAuthor(&Author{UserId: "admin", Password: "admin"})
	fmt.Println(res)
}

func TestMysqlCredentialStorage_GetPasswordByUserId(t *testing.T) {
	s := GetMysqlCredentialStorage(dsn, maxIdle, maxOpen)
	res := s.GetPasswordByUserId("admin")
	fmt.Println(res)
}
