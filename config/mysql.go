package config

type DB struct {
	DriverType      string `mapstructure:"driver-type" json:"driver-type" yaml:"driver-type"`
	Protocol        string `mapstructure:"protocol" json:"protocol" yaml:"protocol"`
	Username        string `mapstructure:"username" json:"username" yaml:"username"`
	Password        string `mapstructure:"password" json:"password" yaml:"password"`
	Host            string `mapstructure:"host" json:"host" yaml:"host"`
	Port            int    `mapstructure:"port" json:"port" yaml:"port"`
	DBName          string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	ConnMaxLifeTime int    `mapstructure:"conn-max-life-tim" json:"conn-max-life-tim" yaml:"conn-max-life-tim"`
	ConnMaxIdleTime int    `mapstructure:"conn-max-idle-time" json:"conn-max-idle-time" yaml:"conn-max-idle-time"`
	MaxOpenConns    int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	MaxIdleConns    int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
}
