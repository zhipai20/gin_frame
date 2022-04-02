package command

import (
	"fmt"
	"kang/bootstrap"
	"kang/config"
	"os"
)

//待完善
// GenerateModel 创建model
func GenerateModel() {
	env := "dev"
	args := os.Args
	if len(args) < 3 {
		fmt.Println("参数缺失：至少需要一个参数 {n} {env}")
		return
	}
	if len(args) >= 4 {
		env = args[3]
	}
	config.ConfEnv = env
	config.InitConfig()
	bootstrap.BootMysql()


}
