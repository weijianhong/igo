package main

import (
	"github.com/weijianhong/igo/core"
	"github.com/weijianhong/igo/global"
	"github.com/weijianhong/igo/initialize"
)

func main() {

	//初始化配置
	global.VP = core.Viper()

	initialize.OtherInit()

	// 初始化log
	global.LOG = core.Zap()

	// 初始化数据库（直接调用驱动的方式）
	initialize.DB()

	// 初始化gorm
	//initialize.Gorm()

	//初始化 redis
	initialize.Redis()

	// 初始化 elasticsearch
	//initialize.ES()

	// 启动grpc服务
	//core.RunGrpcServer()

	// 启动http服务
	core.RunHttpServer()
}
