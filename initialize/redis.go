package initialize

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/weijianhong/igo/global"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	client := &redis.Client{}
	// 单机  standalone  主从 replication 哨兵 sentinel
	switch redisCfg.Deployment {
	case "sentinel", "replication":
		client = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    redisCfg.MasterName,
			Password:      redisCfg.Password,
			DB:            redisCfg.DB,
			SentinelAddrs: redisCfg.SentinelAddrs,
		})

	case "standalone":
		client = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password, // no password set
			DB:       redisCfg.DB,       // use default DB
		})
	default:
		client = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password, // no password set
			DB:       redisCfg.DB,       // use default DB
		})
	}

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.Error(err))
		panic(err)
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
		global.LOG.Info("redis init success ")
	}
}
