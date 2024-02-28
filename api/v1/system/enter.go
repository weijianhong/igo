package system

import "github.com/weijianhong/igo/service"

type ApiGroup struct {
	SystemApi
}

var (
	systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
)
