package command

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	dmysql "github.com/golang-migrate/migrate/v4/database/mysql"

	"github.com/gookit/goutil/dump"

	"kang/config"
)

//未完善
func GenerateMigrate() {
	env := "env"
	args := os.Args
	if len(args) < 4 {
		fmt.Println("参数缺失：至少需要一个参数 {n} {env}")
		return
	}
	n, _ := strconv.Atoi(args[2])
	if len(args) >= 4 {
		env = args[3]
	}
	config.ConfEnv = env
	config.InitConfig()

	mysqlConf := mysql.Config{
		DBName:               config.Conf.Mysql.DbName,
		User:                 config.Conf.Mysql.User,
		Passwd:               config.Conf.Mysql.Kang,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", config.Conf.Mysql.Host, config.Conf.Mysql.Port),
		AllowNativePasswords: true,
		MultiStatements:      true,
		ParseTime:            true,
		Loc:                  time.Local,
	}

	db, err := sql.Open("mysql", mysqlConf.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	driver, err := dmysql.WithInstance(db, &dmysql.Config{
		DatabaseName: config.Conf.Mysql.DbName,
	})
	if err != nil {
		panic(err.Error())
	}

	m, _ := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql",
		driver)

	if err := m.Steps(n); err != nil {
		dump.P("2")
		log.Fatal(err.Error())
	}
}
