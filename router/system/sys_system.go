package system

import (
	"github.com/gin-gonic/gin"
)

type sysRouter struct{}

var sysRouterNew = new(sysRouter)

func (s *sysRouter) Add(Router *gin.RouterGroup) {

	Router.POST("getSystemConfig", systemApi.GetSystemConfig) // 获取配置文件内容

}
