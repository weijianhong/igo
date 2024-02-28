package system

import (
	"github.com/gin-gonic/gin"
	"github.com/weijianhong/igo/model/common/response"
)

type SystemApi struct{}

func (s *SystemApi) GetSystemConfig(c *gin.Context) {
	systemConfigService.AAA()
	response.OkWithDetailed(nil, "获取成功", c)
}
