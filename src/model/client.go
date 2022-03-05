package model

import "encoding/json"

type Client struct {
	// 请求api版本
	Version string `form:"version" json:"version" binding:"required"`
	// 设备平台
	DevicePlatform string `form:"device_platform" json:"device_platform" binding:"required"`
	// 设备 Id
	DeviceId string `form:"device_id" json:"device_id" binding:"required"`
	// 安卓的系统版本
	OsApi int `form:"os_api" json:"os_api" binding:"required"`
	// 渠道
	Channel string `form:"channel" json:"channel" binding:"required"`
	// 应⽤⼤版本，⽐如8.1.4?
	VersionCode string `form:"version_code" json:"version_code" binding:"required"`
	// 应⽤⼩版本，⽐如8.1.4.01?
	UpdateVersionCode string `form:"update_version_code" json:"update_version_code" binding:"required"`
	// app的唯⼀标识?
	Aid int `form:"aid" json:"aid" binding:"required"`
	// 设备的cpu架构: 32/64
	CpuArch string `form:"cpu_arch" json:"cpu_arch" binding:"required"`
}

func (client Client) String() string {
	buf, err := json.Marshal(client)
	if err != nil {
		return ""
	}
	return string(buf)
}

func CreateClientFromString(string string) *Client {
	var client Client
	err := json.Unmarshal([]byte(string), &client)
	if err != nil {
		return nil
	}
	return &client
}
