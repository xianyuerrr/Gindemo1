package operationTest

import (
	"net/url"
)

var ReleaseTest Operation

func init() {
	ReleaseTest = Operation{
		Action: "release",
		FormData: url.Values{
			"id": {"1"},
		},
		NeedAuth: true,
	}
}
