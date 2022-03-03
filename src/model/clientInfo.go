package model

type Client struct {
	// 请求api版本
	Version string `json:"version" binding:"required"`
	// 设备平台
	DevicePlatform string `json:"device_platform" binding:"required"`
	// 设备 ID
	DeviceId string `json:"device_id" binding:"required"`
	// 安卓的系统版本
	OsApi int `json:"os_api" binding:"required"`
	// 渠道
	Channel string `json:"channel" binding:"required"`
	// 应⽤⼤版本，⽐如8.1.4?
	VersionCode string `json:"version_code" binding:"required"`
	// 应⽤⼩版本，⽐如8.1.4.01?
	UpdateVersionCode string `json:"update_version_code" binding:"required"`
	// app的唯⼀标识?
	Aid int `json:"aid" binding:"required"`
	// 设备的cpu架构: 32/64
	CpuArch string `json:"cpu_arch" binding:"required"`
}
