package service

import "demo1/src/model"

func Hit(form model.Client) (string, string, string, string, string) {
	//todo
	rules := model.GetRules(form.DeviceId)

	hitRule := rules[0]
	//if hit,
	//return downloadUrl, updateVersionCode, md5, title, updateTips

	return hitRule.DownloadUrl, hitRule.UpdateVersionCode, hitRule.Md5, hitRule.Title, hitRule.UpdateTips
}
