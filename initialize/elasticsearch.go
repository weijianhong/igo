package initialize

import (
	"fmt"
	"github.com/weijianhong/igo/global"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
)

var (
	esClient *elasticsearch.Client
	once     sync.Once
)

func ES() (*elasticsearch.Client, error) {
	once.Do(func() {
		esCfg := global.CONFIG.Es
		cfg := elasticsearch.Config{
			Addresses: esCfg.Addresses,
			Username:  esCfg.Username,
			Password:  esCfg.Password,
		}

		// 连接Elasticsearch
		var err error
		esClient, err = connectElasticsearch(cfg)
		if err != nil {
			global.LOG.Error("Failed to connect to Elasticsearch", zap.Error(err))
			return
		}

		// 验证连接是否成功
		if err := pingElasticsearch(esClient); err != nil {
			global.LOG.Error("Failed to ping Elasticsearch", zap.Error(err))
			return
		}
	})

	if esClient == nil {
		return nil, fmt.Errorf("Elasticsearch client not initialized")
	}

	return esClient, nil
}

func connectElasticsearch(cfg elasticsearch.Config) (*elasticsearch.Client, error) {
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating the client: %w", err)
	}
	return client, nil
}

func pingElasticsearch(client *elasticsearch.Client) error {
	// 在此添加重试机制，以确保连接稳定
	for i := 0; i < 3; i++ {
		res, err := client.Info()
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		res.Body.Close()
		return nil
	}
	return fmt.Errorf("failed to ping Elasticsearch after 3 attempts")
}
