package config

type ElasticSearch struct {
	Addresses []string `mapstructure:"addresses" json:"addresses" yaml:"addresses"`
	Username  string   `mapstructure:"username" json:"username" yaml:"username"`
	Password  string   `mapstructure:"password" json:"password" yaml:"password"`
}
