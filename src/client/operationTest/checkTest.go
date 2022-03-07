package operationTest

import (
	"net/url"
)

var CheckTest Operation

func init() {
	CheckTest = Operation{
		Action: "check",
		FormData: url.Values{
			"version":             {"8.1.3.0"},
			"device_platform":     {"Android"},
			"device_id":           {"xiaomi11abc"},
			"os_api":              {"9"},
			"channel":             {"xiaomi"},
			"version_code":        {"8.1.2"},
			"update_version_code": {"8.1.2.1"},
			"aid":                 {"15"},
			"cpu_arch":            {"64"},
		},
		NeedAuth: false,
	}
}
