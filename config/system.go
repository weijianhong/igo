package config

type ServerConfig struct {
	HTTP Config `mapstructure:"http" json:"http" yaml:"http"`
	GRPC Config `mapstructure:"grpc" json:"grpc" yaml:"grpc"`
}

type Config struct {
	IsOpen bool   `mapstructure:"is-open" json:"is_open" yaml:"is-open"`
	Addr   string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port   int    `mapstructure:"port" json:"port" yaml:"port"`
}

type System struct {
	Env    string       `mapstructure:"env" json:"env" yaml:"env"`
	Server ServerConfig `mapstructure:"server" json:"server" yaml:"server"`
}
