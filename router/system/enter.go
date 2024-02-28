package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/weijianhong/igo/api/v1"
)

type group struct{}

var GroupNew = new(group)

var (
	//对硬接口
	systemApi = v1.ApiGroupApp.ApiGroup.SystemApi
)

func (a group) Add(Router *gin.RouterGroup) {

	Router.Group("system")

	SysRouterNew.Add(Router)

	return
}
