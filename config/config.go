package config

type Server struct {
	System  System        `mapstructure:"system" json:"system" yaml:"system"`
	Zap     Zap           `mapstructure:"zap" json:"zap" yaml:"zap"`
	DB      DB            `mapstructure:"db" json:"db" yaml:"db"`
	Redis   Redis         `mapstructure:"redis" json:"redis" yaml:"redis"`
	JWT     JWT           `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Es      ElasticSearch `mapstructure:"es" json:"es" yaml:"es"`
	Captcha Captcha       `mapstructure:"captcha" json:"captcha" yaml:"captcha"`

	//Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
