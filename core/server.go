package core

import (
	"app/global"
	"app/initialize"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunHttpServer() {
	httpCfg := global.CONFIG.System.Server.HTTP
	if !httpCfg.IsOpen {
		return
	}

	// 初始化路由
	router := initialize.Routers()

	address := fmt.Sprintf(":%d", httpCfg.Port)
	server := initServer(address, router)

	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))

	// 设置关闭信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		global.LOG.Info("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			global.LOG.Error("Server forced to shutdown:", zap.Error(err))
		}
	}()

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		global.LOG.Error("server run failed", zap.Error(err))
	}
}

func initServer(address string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func RunGrpcServer() {
	grpcCfg := global.CONFIG.System.Server.GRPC
	if !grpcCfg.IsOpen {
		return
	}

	grpcAddr := fmt.Sprintf(":%d", grpcCfg.Port)
	fmt.Println("grpc server start .... ", grpcAddr)

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		global.LOG.Error("grpc failed to listen", zap.Error(err))
		return
	}

	s := grpc.NewServer()
	reflection.Register(s)

	// 注册grpc服务
	initialize.Registers(s)

	go func() {
		if err := s.Serve(lis); err != nil {
			global.LOG.Error("grpc failed to serve", zap.Error(err))
			return
		}
	}()

	// 优雅关闭 gRPC 服务
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down gRPC server...")
	s.GracefulStop()
	fmt.Println("gRPC server shutdown.")
}
