package system

import "github.com/weijianhong/igo/service"

type ApiGroup struct {
	SystemApi
	BaseApi
}

// 对应服务
var (
	systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	jwtService          = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService         = service.ServiceGroupApp.SystemServiceGroup.UserService
)
