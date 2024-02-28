package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/weijianhong/igo/api/v1"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system")
	systemApi := v1.ApiGroupApp.ApiGroup.SystemApi
	{
		sysRouter.POST("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容
	}
}
