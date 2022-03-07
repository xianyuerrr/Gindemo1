package cache

import (
	"grayRelease/src/model"
)

type CheckCache interface {
	Hit(client *model.Client) *string
	Store(client *model.Client, res string) bool
}
