package command

import (
	"fmt"
	"kang/bootstrap"
	"kang/config"
	"kang/global"
	"kang/models"
	util2 "kang/pkg/util"
	"os"
	"time"
)

type GinAdmin models.GinAdmin

// GenerateAdmin 生成admin信息工具
func GenerateAdmin() {
	var env = "dev"
	args := os.Args
	if len(args) < 4 {
		fmt.Println("参数缺失：需要两个参数 {account} {password}")
		return
	}
	if len(args) > 4 {
		env = args[4]
	}
	account := args[2]
	password := args[3]
	config.ConfEnv = env
	config.InitConfig()
	bootstrap.BootMysql()

	salt := util2.GenerateUuid(32)

	user := GinAdmin{
		Uuid:         util2.GenerateBaseSnowId(32),
		Account:      account,
		Password:     util2.GeneratePasswordHash(password, salt),
		Phone:        "12345678901",
		Avatar:       "",
		Salt:         salt,
		RealName:     account,
		RegisterTime: uint64(time.Now().Unix()),
		RegisterIp:   "127.0.0.1",
		LoginTime:    uint64(time.Now().Unix()),
		LoginIp:      "127.0.0.1",
		Status:       1,
		CreatedAt:    uint64(time.Now().Unix()),
		UpdatedAt:    uint64(time.Now().Unix()),
	}

	err := global.DB.Create(&user).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("账号：" + account + "; 密码：" + password + " 生成成功")
	return
}
