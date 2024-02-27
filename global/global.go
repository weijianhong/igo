package global

import (
	"app/config"
	"database/sql"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	CONFIG config.Server

	VP    *viper.Viper
	LOG   *zap.Logger
	DB    *sql.DB
	REDIS *redis.Client
	ES    *elasticsearch.Client
)
