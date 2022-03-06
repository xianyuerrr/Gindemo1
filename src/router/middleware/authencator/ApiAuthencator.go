package authencator

type ApiAuthenticator interface {
	Auth(apiRequest ApiRequest) bool
}
