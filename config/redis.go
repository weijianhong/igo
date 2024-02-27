package config

type Redis struct {
	Deployment    string   `mapstructure:"deployment" json:"deployment" yaml:"deployment"`             //
	Addr          string   `mapstructure:"addr" json:"addr" yaml:"addr"`                               // 服务器地址:端口
	Password      string   `mapstructure:"password" json:"password" yaml:"password"`                   // 密码
	DB            int      `mapstructure:"db" json:"db" yaml:"db"`                                     // redis的哪个数据库
	MasterName    string   `mapstructure:"master-name" json:"master-name" yaml:"master-name"`          // redis的哪个数据库
	SentinelAddrs []string `mapstructure:"sentinel-addrs" json:"sentinel-addrs" yaml:"sentinel-addrs"` // redis的哪个数据库
}
