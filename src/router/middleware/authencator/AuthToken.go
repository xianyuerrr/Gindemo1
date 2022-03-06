package authencator

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

type AuthToken struct {
	token               string
	createTime          int64
	expiredTimeInterval int64
}

func generateToken(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func CreateAuthToken(baseUrl string, user string, password string, createTime int64) *AuthToken {
	token := generateToken(baseUrl + user + password + strconv.FormatInt(createTime, 10))
	return &AuthToken{token: token, createTime: createTime, expiredTimeInterval: int64(time.Minute.Seconds() * 30)}
}

func (authToken AuthToken) GetToken() string {
	return authToken.token
}

func (authToken AuthToken) IsExpired() bool {
	now := time.Now().Unix()
	return !(authToken.createTime <= now && now <= authToken.createTime+authToken.expiredTimeInterval)
}

func (authToken AuthToken) Match(another AuthToken) bool {
	return authToken.GetToken() == another.GetToken()
}
