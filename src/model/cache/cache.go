package cache

import (
	"grayRelease/src/model"
)

type Cache interface {
	Hit(client *model.Client) *string
	Store(client *model.Client, res string) bool
}
