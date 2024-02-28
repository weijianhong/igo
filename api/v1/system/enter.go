package system

import "github.com/weijianhong/igo/service"

type ApiGroup struct {
	SystemApi
}

// 对应服务
var (
	systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
)
