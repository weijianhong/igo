package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/weijianhong/igo/api/v1"
)

type group struct{}

var GroupNew = new(group)

var (
	//对硬接口
	systemApi = v1.ApiGroupApp.SystemApiGroup.SystemApi
)

func (a group) Add(Router *gin.RouterGroup) {

	Router.Group("system")

	sysRouterNew.Add(Router)

	return
}
