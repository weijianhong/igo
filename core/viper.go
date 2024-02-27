package core

import (
	"app/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"sync"
)

var (
	v     *viper.Viper
	onceV sync.Once
)

// Viper 初始化 Viper 实例
func Viper(path ...string) *viper.Viper {
	onceV.Do(func() {
		var config string

		flag.StringVar(&config, "c", "config.yaml", "choose config file.")
		flag.Parse()

		v = viper.New()
		v.SetConfigFile(config)
		v.SetConfigType("yaml")
		err := v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		v.WatchConfig()

		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file changed:", e.Name)
			if err = v.Unmarshal(&global.CONFIG); err != nil {
				fmt.Println(err)
			}
		})
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			panic(err)
		}
	})

	return v
}
