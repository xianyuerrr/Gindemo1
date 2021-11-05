package model

type Client struct {
	Version           string `json:"version" binding:"required"`
	DevicePlatform    string `json:"device_platform" binding:"required"`
	DeviceId          string `json:"device_id" binding:"required"`
	OsApi             string `json:"os_api" binding:"required"`
	Channel           string `json:"channel" binding:"required"`
	VersionCode       string `json:"version_code" binding:"required"`
	UpdateVersionCode string `json:"update_version_code" binding:"required"`
	Aid               int    `json:"aid" binding:"required"`
	CpuArch           int    `json:"cpu_arch" binding:"required"`
}
