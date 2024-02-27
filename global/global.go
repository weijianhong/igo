package global

import (
	"app/config"
	"database/sql"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
)

var (
	CONFIG config.Server

	VP    *viper.Viper
	LOG   *zap.Logger
	DB    *sql.DB
	REDIS *redis.Client
	ES    *elasticsearch.Client

	BlackCache local_cache.Cache

	GVA_Concurrency_Control = &singleflight.Group{}
)
