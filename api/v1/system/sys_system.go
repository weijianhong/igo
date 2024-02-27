package system

import (
	"app/model/common/response"
	"github.com/gin-gonic/gin"
)

type SystemApi struct{}

func (s *SystemApi) GetSystemConfig(c *gin.Context) {
	response.OkWithDetailed(nil, "获取成功", c)
}
