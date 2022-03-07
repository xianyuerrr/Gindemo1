package operationTest

import (
	"net/url"
)

var DeleteTest Operation

func init() {
	DeleteTest = Operation{
		Action: "delete",
		FormData: url.Values{
			"id": {"2"},
		},
		NeedAuth: true,
	}
}
