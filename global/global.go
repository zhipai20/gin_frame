package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"kang/config"
)

var (
	G_DB     *gorm.DB       // Mysql数据库
	G_Logger *zap.Logger    // 日志
	G_Redis  *redis.Client  // redis连接池
	G_Conf   *config.App //配置

	ConfEnv string //环境
)
