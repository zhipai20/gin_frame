package config

type App struct {
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
	Log    Log    `mapstructure:"log" json:"log" yaml:"log"`

	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`

	Oss   Oss   `mapstructure:"oss" json:"oss" yaml:"oss"`
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
