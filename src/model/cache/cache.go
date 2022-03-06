package cache

import (
	"grayRelease/src/model/tables"
)

type Cache interface {
	Hit(client *tables.Client) *string
	Store(client *tables.Client, res string) bool
}
