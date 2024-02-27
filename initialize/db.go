package initialize

import (
	"app/global"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"sync"
	"time"
)

var (
	db     *sql.DB
	onceDB sync.Once
)

func DB() {
	onceDB.Do(func() {
		cfg := global.CONFIG.DB

		var err error
		db, err = sql.Open(cfg.DriverType, fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
			cfg.Username, cfg.Password, cfg.Protocol, cfg.Host, cfg.Port, cfg.DBName))
		if err != nil {
			global.LOG.Error("Error connecting to database", zap.Error(err))
			return
		}

		db.SetMaxOpenConns(cfg.MaxOpenConns)
		db.SetMaxIdleConns(cfg.MaxIdleConns)
		db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime))
		db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifeTime))
	})

	global.DB = db
}
