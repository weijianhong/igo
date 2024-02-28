package core

import (
	"fmt"
	"github.com/weijianhong/igo/core/internal"
	"github.com/weijianhong/igo/global"
	"github.com/weijianhong/igo/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	logger  *zap.Logger
	onceZap sync.Once
)

func Zap() *zap.Logger {
	onceZap.Do(func() {
		if ok, _ := utils.PathExists(global.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
			fmt.Printf("create %v directory\n", global.CONFIG.Zap.Director)
			_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
		}

		cores := internal.Zap.GetZapCores()
		logger = zap.New(zapcore.NewTee(cores...))

		if global.CONFIG.Zap.ShowLine {
			logger = logger.WithOptions(zap.AddCaller())
		}

	})
	return logger
}
