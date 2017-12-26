package main

import (
	"fmt"
	"github.com/pippozq/netadmin/models"
	_ "github.com/pippozq/netadmin/routers"
	"github.com/pippozq/netadmin/schedules"
	"github.com/pippozq/netadmin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/lib/pq"
)

func init() {
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

	orm.RegisterDataBase("default", "postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		beego.AppConfig.String("user"),
		beego.AppConfig.String("password"),
		beego.AppConfig.String("db_name"),
		beego.AppConfig.String("host"),
		beego.AppConfig.String("port")), maxIdleConn, maxOpenConn)

	orm.RunSyncdb("default", false, true)

	initUserDone := true
	initUserDone = models.InitUser()
	if !initUserDone {
		beego.Error("Init User Error")
	}
}

func main() {
	beego.SetLevel(beego.LevelInformational)

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.BConfig.WebConfig.StaticDir["/"] = "static"
	beego.BConfig.WebConfig.StaticDir["/assets"] = "static/assets"
	beego.BConfig.WebConfig.StaticDir["/pages"] = "static/pages"

	schedules.InitTask()
	beego.ErrorController(&utils.BaseController{})
	beego.Run()

}
