package config

type Server struct {
	Mode            string `yaml:"mode"`
	DefaultPageSize int    `yaml:"defaultPageSize"`
	MaxPageSize     int    `yaml:"maxPageSize"`
	TokenExpire     int64  `yaml:"tokenExpire"`
	TokenKey        string `yaml:"tokenKey"`
	TokenIssuer     string `yaml:"tokenIssuer"`
	JwtSecret       string `yaml:"jwtSecret"`
}
