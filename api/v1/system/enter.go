package system

import "app/service"

type ApiGroup struct {
	SystemApi
}

var (
	systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
)
