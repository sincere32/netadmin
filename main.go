package main

import (
	"flag"
	"fmt"
	_ "gitee.com/pippozq/netadmin/routers"
	"gitee.com/pippozq/netadmin/schedules"
	"gitee.com/pippozq/netadmin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/lib/pq"
	"gitee.com/pippozq/netadmin/models"
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
		orm.RunSyncdb("default", false, true)
	}

	return err
}

func main() {
	beego.SetLevel(beego.LevelInformational)

	initDb := flag.String("sync_db", "no", "init db")
	initUser := flag.String("init_user", "admin", "Init Admin User")
	flag.Parse()

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	err := InitDBPool(*initDb)
	if err != nil {
		beego.Error(fmt.Sprintf("Mode:%s, Init DB Error:%s", *initDb, err))
	} else {
		schedules.InitTask()

		if *initUser == "yes"{
			models.InitUser()
		}

		beego.ErrorController(&utils.BaseController{})
		beego.Run()
	}
}
