package operationTest

import (
	"net/url"
)

var AddTest Operation

func init() {
	AddTest = Operation{
		Action: "add",
		FormData: url.Values{
			"id":                      {"2"},
			"aid":                     {"15"},
			"platform":                {"Android"},
			"download_url":            {"download_url.apk"},
			"update_version_code":     {"8.1.3.03"},
			"md_5":                    {"md_5"},
			"device_id_list":          {"xiaomi10, xiaomi11"},
			"max_update_version_code": {"8.1.3.0"},
			"min_update_version_code": {"8.0.3.0"},
			"max_os_api":              {"9"},
			"min_os_api":              {"6"},
			"cpu_arch":                {"64"},
			"channel":                 {"xiaomi"},
			"title":                   {"a title"},
			"update_tips":             {"a update_tips"},
		},
		NeedAuth: true,
	}
}
