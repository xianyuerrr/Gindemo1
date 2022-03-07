package operationTest

import "net/url"

type Operation struct {
	Action   string
	FormData url.Values
	NeedAuth bool
}
