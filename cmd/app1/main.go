package main

import (
	"app/core"
	"app/global"
	"app/initialize"
)

func main() {

	//初始化配置
	global.VP = core.Viper()
	//fmt.Println(global.CONFIG)

	// 初始化log
	global.LOG = core.Zap()

	// 初始化数据库（直接调用驱动的方式）
	initialize.DB()

	//初始化 redis
	//initialize.Redis()

	// 初始化 elasticsearch
	//initialize.ES()

	// 启动grpc服务
	core.RunGrpcServer()

	// 启动http服务
	core.RunHttpServer()

}
