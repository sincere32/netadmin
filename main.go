package main

import (
	"flag"
	"fmt"
	_ "gitee.com/pippozq/netadmin/models"
	_ "gitee.com/pippozq/netadmin/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"gitee.com/pippozq/netadmin/utils"
)

func InitDBPool(initDb string) error {
	orm.Debug, _ = beego.AppConfig.Bool("debug")
	orm.RegisterDriver("postgres", orm.DRPostgres)
	maxIdleConn, err := beego.AppConfig.Int("max_idle_conn")
	if err != nil {
		maxIdleConn = 10
	}

	maxOpenConn, err := beego.AppConfig.Int("max_open_conn")

	if err != nil {
		maxOpenConn = 10
	}

	err = orm.RegisterDataBase("default", "postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		beego.AppConfig.String("user"),
		beego.AppConfig.String("password"),
		beego.AppConfig.String("db_name"),
		beego.AppConfig.String("host"),
		beego.AppConfig.String("port")), maxIdleConn, maxOpenConn)

	if err == nil && initDb == "yes" {
		orm.RunSyncdb("default", true, true)
	}

	return err
}
func main() {
	beego.SetLevel(beego.LevelInformational)

	initDb := flag.String("syncdb", "no", "init db")
	mode := flag.String("mode", "test", "run mode")
	flag.Parse()

	beego.BConfig.RunMode = *mode

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	err := InitDBPool(*initDb)
	if err != nil {
		beego.Error(fmt.Sprintf("Mode:%s, Init DB Error:%s", *initDb, err))
	} else {
		beego.ErrorController(&utils.BaseController{})
		beego.Run()
	}
}
