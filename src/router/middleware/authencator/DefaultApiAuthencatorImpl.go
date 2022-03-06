package authencator

import (
	"fmt"
	"time"
)

type DefaultApiAuthencatorImpl struct {
	credentialStorage CredentialStorage
}

func GeDefaulttApiAuthencator(credentialStorage CredentialStorage) *DefaultApiAuthencatorImpl {
	if credentialStorage == nil {
		return nil
	}
	return &DefaultApiAuthencatorImpl{credentialStorage: credentialStorage}
}

func (d DefaultApiAuthencatorImpl) Auth(apiRequest ApiRequest) bool {
	token := apiRequest.getToken()
	timestamp := apiRequest.getTimeStamp()

	clientAuthToken := AuthToken{token: token, createTime: timestamp, expiredTimeInterval: int64(time.Minute.Seconds() * 30)}

	if clientAuthToken.IsExpired() {
		fmt.Println("Token is expired.")
		return false
	}

	baseUrl := apiRequest.getBaseUrl()
	appId := apiRequest.getAppId()

	password := d.credentialStorage.GetPasswordByUserId(appId)
	serverAuthToken := CreateAuthToken(baseUrl, appId, password, timestamp)
	if !serverAuthToken.Match(clientAuthToken) {
		fmt.Println("Token verification failed.")
		return false
	}
	return true
}
