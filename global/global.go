package global

import (
	"database/sql"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"github.com/weijianhong/igo/config"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	CONFIG config.Server

	VP    *viper.Viper
	LOG   *zap.Logger
	DB    *sql.DB
	REDIS *redis.Client
	ES    *elasticsearch.Client

	BlackCache local_cache.Cache

	ConcurrencyControl = &singleflight.Group{}

	DBG *gorm.DB
)
