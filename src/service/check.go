package service

import (
	"github.com/spf13/cast"
	"grayRelease/src/model"
	"grayRelease/src/model/cache"
	"math/rand"
	"strings"
)

func Hit(client *model.Client) []string {
	// 查询缓存
	var checkCache cache.CheckCache
	checkCache = cache.GetCheckRedisCache()
	resCache := checkCache.Hit(client)
	if resCache != nil {
		return strings.Split(*resCache, ",")
	}
	// if hit,
	// return []string {downloadUrl, updateVersionCode, md5, title, updateTips}

	// 新版本检查接⼝规则，多条件的⽐较顺序是：
	// 业务id > platform > 渠道 > 设备⽩名单 > 【其他条件计算顺序均可】
	// 如果命中，返回满⾜条件的升级包的基本信息；⾄多只能返回⼀条升级包规则；

	// 根据 业务id > platform > 渠道 获取符合条件的 rules
	rules := GetReleasedRules(client.Aid, client.DevicePlatform, client.Channel)
	if rules == nil || cap(rules) == 0 {
		return getDownloadInfo(nil)
	}
	// order by updateVersionCode, use quickSort O(nlogn)
	quickSort(&rules, 0, len(rules)-1)

	// 对于 rule, 显然 rule.UpdateVersionCode > rule.MaxUpdateVersionCode >= rule.MinUpdateVersionCode
	// 按照常理来说，对 rule.UpdateVersionCode 进行排序之后，rule.MaxUpdateVersionCode 和 rule.MinUpdateVersionCode 也是有序的
	// 所以进行二分，找到 client.UpdateVersionCode 符合的 rule 子集
	left, right := binarySearchLeft(rules, client), binarySearchRight(rules, client)

	// 由于 rules 是按照从小到大的顺序排列，而有多个匹配的规则时，应该返回 UpdateVersion 最大的（即返回最新版本的安装包链接）
	// 所以需要 从 right 到 left 进行匹配，取第一个命中的 rule
	var hitRule model.Rule
	for i := right; i >= left; i-- {
		if matchRule(&(rules)[i], client) {
			hitRule = rules[i]
			break
		}
	}
	res := getDownloadInfo(&hitRule)
	checkCache.Store(client, strings.Join(res, ","))
	return res
}

func quickSort(rules *[]model.Rule, l, r int) {
	if l >= r {
		return
	}

	randIdx := rand.Intn(r-l+1) + l
	(*rules)[l], (*rules)[randIdx] = (*rules)[randIdx], (*rules)[l]
	mid := l
	target := (*rules)[l].UpdateVersionCode
	for i, rule := range (*rules)[l : r+1] {
		if rule.UpdateVersionCode <= target {
			(*rules)[mid], (*rules)[l+i] = (*rules)[l+i], (*rules)[mid]
			mid++
		}
	}

	(*rules)[mid-1], (*rules)[l] = (*rules)[l], (*rules)[mid-1]

	// 左侧 小于等于 target
	// 右侧 大于 target
	quickSort(rules, l, mid-2)
	quickSort(rules, mid, r)
	return
}

func matchRule(rule *model.Rule, client *model.Client) bool {
	// 判断 deviceId 是否在 rule 的白名单
	if ExistsDeviceIdInWhiteList(int(rule.Id), client.DeviceId) {
		return true
	}

	// model.Client.Version : 请求api版本，⽐如v1/v2
	// model.Client.version_code : 应⽤⼤版本，⽐如8.1.4
	// 是否符合 版本要求（应⽤⼩版本，⽐如8.1.4.01），将版本筛选放到前面了，此函数不再需要负责此部分工作了
	if client.OsApi < rule.MinOsApi || client.OsApi > rule.MaxOsApi {
		// 系统 是否适配
		return false
	}
	if client.CpuArch != rule.CpuArch {
		return false
	}
	return true
}

func binarySearchRight(rules []model.Rule, client *model.Client) int {
	l, r := 0, len(rules)-1
	for l < r {
		m := l + ((r-l)>>1 + 1) // 很怪，>> 比 r-l 先执行？？？
		// if client.UpdateVersionCode < (*rules)[m].MinUpdateVersionCode
		if compareUpdateVersionCode(client.UpdateVersionCode, rules[m].MinUpdateVersionCode) == -1 {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func binarySearchLeft(rules []model.Rule, client *model.Client) int {
	l, r := 0, len(rules)-1
	for l < r {
		m := l + ((r - l) >> 1) // 很怪，>> 比 r-l 先执行？？？
		// if client.UpdateVersionCode > (*rules)[m].MaxUpdateVersionCode {
		if compareUpdateVersionCode(client.UpdateVersionCode, (rules)[m].MaxUpdateVersionCode) == 1 {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func compareUpdateVersionCode(UpdateVersionCode1, UpdateVersionCode2 string) int {
	/*
		比较 UpdateVersionCode 大小的函数，UpdateVersionCode1 和 UpdateVersionCode2 格式相同(eg. "1.1.1.1")
		返回值解析 :
		-1 : UpdateVersionCode1 < UpdateVersionCode2
		0 : UpdateVersionCode1 == UpdateVersionCode2
		1 : UpdateVersionCode1 > UpdateVersionCode2
	*/
	var res = 0
	lis1 := strings.Split(UpdateVersionCode1, ".")
	lis2 := strings.Split(UpdateVersionCode2, ".")
	for index := 0; index < len(lis1); index++ {
		if cast.ToInt(lis1[index]) < cast.ToInt(lis2[index]) {
			res = -1
			break
		}
		if cast.ToInt(lis1[index]) > cast.ToInt(lis2[index]) {
			res = 1
			break
		}
	}
	return res
}

func getDownloadInfo(rule *model.Rule) []string {
	if nil == rule {
		return []string{"", "", "", "", ""}
	}
	return []string{(*rule).DownloadUrl, (*rule).UpdateTips, (*rule).Md5, (*rule).Title, (*rule).UpdateTips}
}
