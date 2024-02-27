package system

import (
	v1 "app/api/v1"
	"github.com/gin-gonic/gin"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system")
	systemApi := v1.ApiGroupApp.ApiGroup.SystemApi
	{
		sysRouter.POST("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容
	}
}
