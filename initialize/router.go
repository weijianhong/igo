package initialize

import (
	"fmt"
	"github.com/weijianhong/igo/router"
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weijianhong/igo/global"
	"github.com/weijianhong/igo/middleware"
)

// 初始化总路由

func Routers() *gin.Engine {

	// 设置为发布模式
	if global.CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode) //DebugMode ReleaseMode TestMode
	}

	Router := gin.New()
	Router.Use(gin.Recovery())
	if global.CONFIG.System.Env != "public" {
		Router.Use(gin.Logger())
	}

	serverRunDir := "."
	global.LOG.Info("", zap.String("serverRunDir", serverRunDir))

	// 静态 路由功能
	Router.LoadHTMLGlob(fmt.Sprintf("%s/dist/*.html", serverRunDir)) // npm打包成dist的路径
	Router.StaticFile("/favicon.ico", fmt.Sprintf("%s/dist/favicon.ico", serverRunDir))
	Router.Static("/assets", fmt.Sprintf("%s/dist/assets", serverRunDir))              // 添加资源路径
	Router.StaticFile("/", fmt.Sprintf("%s/dist/index.html", serverRunDir))            // 前端网页入口页面
	Router.StaticFS(global.CONFIG.Local.Path, http.Dir(global.CONFIG.Local.StorePath)) // 为用户头像和文件提供静态地址

	// 方便统一添加路由组前缀 多服务器上线使用
	var WsGroup = Router.Group("ws")

	// 不做权限检查的接口
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	// 需要做权限验证的接口
	PrivateGroup := Router.Group("api")
	PrivateGroup.Use(middleware.JWTAuth())

	// 以下为业务路基设置
	router.GroupNew.Add(PublicGroup, PrivateGroup, WsGroup)

	global.LOG.Info("router register success")
	return Router
}
