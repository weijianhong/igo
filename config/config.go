package config

type Server struct {
	System  System        `mapstructure:"system" json:"system" yaml:"system"`
	Zap     Zap           `mapstructure:"zap" json:"zap" yaml:"zap"`
	DB      DB            `mapstructure:"db" json:"db" yaml:"db"`
	Redis   Redis         `mapstructure:"redis" json:"redis" yaml:"redis"`
	JWT     JWT           `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Es      ElasticSearch `mapstructure:"es" json:"es" yaml:"es"`
	Captcha Captcha       `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Cors    CORS          `mapstructure:"cors" json:"cors" yaml:"cors"`
	Local   Local         `mapstructure:"local" json:"local" yaml:"local"`

	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  Pgsql  `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Sqlite Sqlite `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
}
