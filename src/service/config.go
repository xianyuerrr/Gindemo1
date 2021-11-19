package service

import (
	"demo1/src/model"
	"log"
	"strconv"
	"strings"
)

func CheckConfig(config model.Rule) bool {

	//return true
	if !isValidAid(config.Aid) {
		log.Println(" invalid Aid")
		return false
	}
	if !isValidPlatform(config.Platform) {
		log.Println(" invalid Platform")
		return false
	}
	if !isValidDownloadUrl(config.DownloadUrl) {
		log.Println(" invalid DownloadUrl")
		return false
	}
	if !isValidUVC(config.UpdateVersionCode) {
		log.Println(" invalid UpdateVersionCode")
		return false
	}
	if !isValidMd5(config.Md5) {
		log.Println(" invalid Md5")
		return false
	}
	if !isValidDeviceIdList(config.DeviceIdList) {
		log.Println(" invalid DeviceIdList")
		return false
	}
	if !isValidMaxMinUVC(config.MaxUpdateVersionCode, config.MinUpdateVersionCode) {
		log.Println(" invalid MaxUpdateVersionCode or MinUpdateVersionCode")
		return false
	}
	if !isValidMaxMinOsApi(config.MaxOsApi, config.MinOsApi) {
		log.Println(" invalid MaxOsApi or MinOsApi")
		return false
	}
	if !isValidCpuArch(config.CpuArch) {
		log.Println(" invalid CpuArch")
		return false
	}
	if !isValidTitle(config.Title) {
		log.Println(" invalid Title")
		return false
	}
	if !isValidUpdateTips(config.UpdateTips) {
		log.Println(" invalid UpdateTips")
		return false
	}

	//版本号化简
	if MatVersion(config.UpdateVersionCode) {
		config.UpdateVersionCode = simplifyVersionCode(config.UpdateVersionCode)
	} else {
		log.Println(" invalid UpdateVersionCode")
		return false
	}
	if MatVersion(config.MaxUpdateVersionCode) {
		config.MaxUpdateVersionCode = simplifyVersionCode(config.MaxUpdateVersionCode)
	} else {
		log.Println(" invalid MaxUpdateVersionCode")
		return false
	}
	if MatVersion(config.MinUpdateVersionCode) {
		config.MinUpdateVersionCode = simplifyVersionCode(config.MinUpdateVersionCode)
	} else {
		log.Println(" invalid MinUpdateVersionCode")
		return false
	}

	return true
}

func isValidAid(aid int) bool {
	if aid < 0 {
		return false
	}
	return true
}

func isValidPlatform(platform string) bool {
	if platform != "Android" && platform != "iOS" &&
		platform != "HarmonyOS" {
		return false
	}
	return true
}

func isValidDownloadUrl(url string) bool {
	if !strings.HasSuffix(url, ".apk") &&
		!strings.HasSuffix(url, ".ipa") && !strings.HasSuffix(url, ".pxl") && !strings.HasSuffix(url, ".deb") &&
		!strings.HasSuffix(url, ".sis") && !strings.HasSuffix(url, ".sisx") && !strings.HasSuffix(url, ".jar") &&
		!strings.HasSuffix(url, ".xap") {
		return false
	}
	return true
}

func isValidUVC(uvc string) bool {
	nums := strings.Split(uvc, ".")
	for i := 0; i < len(nums); i++ {
		num, err := strconv.Atoi(nums[i])
		if err != nil || num < 0 {
			return false
		}
	}
	return true
}

func isValidMd5(md5 string) bool {
	return true
}

func isValidDeviceIdList(deviceIdList string) bool {
	return true
}

func isValidMaxMinUVC(maxUVC string, minUVC string) bool {
	if !isValidUVC(maxUVC) {
		return false
	}
	if !isValidUVC(minUVC) {
		return false
	}
	if maxUVC < minUVC {
		return false
	}
	return true
}

func isValidMaxMinOsApi(maxOsApi int, minOsApi int) bool {
	if maxOsApi < 0 || minOsApi < 0 || maxOsApi < minOsApi {
		return false
	}
	return true
}

func isValidCpuArch(cpuArch string) bool {
	return true
}

func isValidChannel(channel string) bool {
	return true
}

func isValidTitle(title string) bool {
	return true
}

func isValidUpdateTips(updateTips string) bool {
	return true
}

func MatVersion(VersionCode string) bool {
	temp := strings.Split(VersionCode, ".")
	if len(temp) > 4 {
		return false
	}
	for i := 0; i < len(temp); i++ {
		if len(temp[i]) > 4 {
			return false
		}
	}
	return true
}

func simplifyVersionCode(VersionCode string) string {
	temp := strings.Split(VersionCode, ".")
	var ans string
	for i := 0; i < len(temp); i++ {
		t := -1
		if len(temp[i]) == 0 {
			ans = ans + "0"
		} else {
			for j := 0; j < len(temp[i]); j++ {
				ch := temp[i][j]
				if ch != '0' {
					t = j
					break
				}
			}
			if t != -1 {
				ans = ans + temp[i][t:]
			} else {
				ans = ans + "0"
			}
		}
		if i != len(temp)-1 {
			ans = ans + "."
		}
	}

	if len(temp) < 4 {
		for i := len(temp); i < 4; i++ {
			ans = ans + ".0"
		}
	}
	return ans
}
