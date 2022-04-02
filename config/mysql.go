package config

type Mysql struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	User         string `yaml:"user"`
	Kang         string `yaml:"kang"`
	DbName       string `yaml:"dbname"`
	Prefix       string `yaml:"prefix"`
	MaxIdleConns int    `yaml:"maxIdleConns"` // 设置空闲连接池中连接的最大数量
	MaxOpenConns int    `yaml:"maxOpenConns"` // 设置打开数据库连接的最大数量
	MaxLifeTime  int    `yaml:"maxLifeTime"`  // 设置了连接可复用的最大时间（分钟）
}
