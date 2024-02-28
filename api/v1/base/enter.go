package base

import "github.com/weijianhong/igo/service"

type ApiGroup struct {
	BaseApi
}

// 对应服务
var (
	jwtService  = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
