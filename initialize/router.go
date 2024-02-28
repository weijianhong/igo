package initialize

import (
	"fmt"
	"github.com/weijianhong/igo/utils"
	"go.uber.org/zap"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weijianhong/igo/global"
	"github.com/weijianhong/igo/middleware"
	"github.com/weijianhong/igo/router"
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

	systemRouter := router.RouterGroupApp.System

	serverRunDir, _ := utils.GetCurrentDir()
	serverRunDir = "."
	global.LOG.Info("", zap.String("serverRunDir", serverRunDir))

	// 静态 路由功能
	Router.LoadHTMLGlob(fmt.Sprintf("%s/dist/*.html", serverRunDir)) // npm打包成dist的路径
	Router.StaticFile("/favicon.ico", fmt.Sprintf("%s/dist/favicon.ico", serverRunDir))
	Router.Static("/assets", fmt.Sprintf("%s/dist/assets", serverRunDir))   // 添加资源路径
	Router.StaticFile("/", fmt.Sprintf("%s/dist/index.html", serverRunDir)) // 前端网页入口页面
	// 为用户头像和文件提供静态地址
	Router.StaticFS(global.CONFIG.Local.Path, http.Dir(global.CONFIG.Local.StorePath))

	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	//global.GVA_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		//systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		//systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth())
	{
		//systemRouter.InitApiRouter(PrivateGroup, PublicGroup)    // 注册功能api路由
		systemRouter.InitSystemRouter(PrivateGroup) // system相关路由

	}

	global.LOG.Info("router register success")
	return Router
}
