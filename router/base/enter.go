package base

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/weijianhong/igo/api/v1"
)

type group struct{}

var GroupNew = new(group)

var (
	//对硬接口
	baseApi = v1.ApiGroupApp.ApiSystemGroup.BaseApi
)

func (a group) Add(rg *gin.RouterGroup) {

	rg.Group("base")

	baseRouterNew.Add(rg)

	return
}
