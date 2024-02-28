package base

import (
	"github.com/gin-gonic/gin"
)

type baseRouter struct{}

var baseRouterNew = new(baseRouter)

func (s *baseRouter) Add(rg *gin.RouterGroup) {

	rg.GET("captcha", baseApi.Captcha) // 获取验证码

	rg.POST("login", baseApi.Login)       // 登录
	rg.POST("register", baseApi.Register) // 注册

}
