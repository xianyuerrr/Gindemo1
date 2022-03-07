package operationTest

import (
	"net/url"
)

var OfflineTest Operation

func init() {
	OfflineTest = Operation{
		Action: "offline",
		FormData: url.Values{
			"id": {"2"},
		},
		NeedAuth: true,
	}
}
