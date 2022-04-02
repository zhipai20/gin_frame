package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"kang/global"
	"kang/pkg/lib"
	util2 "kang/pkg/util"
)

// 定义服务列表
const (
	LoggerService = `Logger`
	MysqlService  = `Mysql`
	RedisService  = `Redis`
)

type bootServiceMap map[string]func() error

var (
	err           error
	BootedService []string
	// serviceMap 程序启动时需要自动加载的服务
	serviceMap = bootServiceMap{
		LoggerService: BootLogger,
		MysqlService:  BootMysql,
		RedisService:  BootRedis,
	}
)

// BootService 加载服务
func BootService(services ...string) {
	serviceMap[LoggerService] = BootLogger
	if global.G_Logger != nil {
		global.G_Logger.Info("服务列表已加载完成")
	}

	if len(services) == 0 {
		services = serviceMap.keys()
	}

	BootedService = make([]string, 0)

	//隐藏bug,map的循环输出顺序是不确定的。但却要保证先初始化logger模块
	for k, val := range serviceMap {
		if util2.InStringSlice(k, services) {
			if err := val(); err != nil {
				panic("程序服务启动失败:" + err.Error())
			}
			BootedService = append(BootedService, k)
		}
	}
}

// BootLogger 将配置载入日志服务
func BootLogger() error {
	if global.G_Logger != nil {
		return nil
	}

	global.G_Logger = lib.NewLogger()
	global.G_Logger.Info(fmt.Sprintf("程序载入Logger服务成功 环境名为：%s 日志路径为: %s",global.ConfEnv, global.G_Conf.Log.Director))
	return nil
}

// BootMysql 装配数据库连接
func BootMysql() error {
	if global.G_DB != nil {
		return nil
	}

	dbConfig := lib.DatabaseConfig{
		Host:         global.G_Conf.Mysql.Host,
		Port:         global.G_Conf.Mysql.Port,
		User:         global.G_Conf.Mysql.User,
		Pass:         global.G_Conf.Mysql.Kang,
		DbName:       global.G_Conf.Mysql.DbName,
		Prefix:       global.G_Conf.Mysql.Prefix,
		MaxIdleConns: global.G_Conf.Mysql.MaxIdleConns,
		MaxOpenConns: global.G_Conf.Mysql.MaxOpenConns,
		MaxLifeTime:  global.G_Conf.Mysql.MaxLifeTime,
	}

	global.G_DB, err = lib.NewMysql(dbConfig)
	if err == nil {
			global.G_Logger.Info(":程序载入MySQL服务成功")
	}
	return err

}

// BootRedis 装配redis服务
func BootRedis() error {
	redisConfig := lib.RedisConfig{
		Addr:     fmt.Sprintf("%s:%s", global.G_Conf.Redis.Host, global.G_Conf.Redis.Port),
		Password: global.G_Conf.Redis.Password,
		DbNum:    global.G_Conf.Redis.DbNum,
	}
	global.G_Redis, err = lib.NewRedis(redisConfig)
	if err != nil {
		global.G_Logger.Error("程序载入Redis服务错误：" + err.Error())
	}

	global.G_Logger.Info("程序载入Redis服务成功",)
	return err
}

func BootConfig() {
	var configFile = fmt.Sprintf("config.%s.yaml", global.ConfEnv)

	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.G_Conf); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.G_Conf); err != nil {
		fmt.Println(err)
	}
}

// keys 获取BootServiceMap中所有键值
func (m bootServiceMap) keys() []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
