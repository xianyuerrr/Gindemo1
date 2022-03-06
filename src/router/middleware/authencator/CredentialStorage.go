package authencator

type CredentialStorage interface {
	GetPasswordByUserId(appId string) string
}
