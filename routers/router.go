// @APIVersion 1.0.0
// @Title NetAdmin API
// @Description Netadmin has every tool to get any job done, so codename for the new netadmin APIs.
// @Contact zengqi0529@163.com
package routers

import (
	"github.com/pippozq/netadmin/controllers"
	"github.com/pippozq/netadmin/filters"
	"github.com/astaxie/beego"
)

func init() {

	beego.InsertFilter("/*", beego.BeforeExec, filters.LoginFilter)
	beego.InsertFilter("/*", beego.BeforeExec, filters.RoleFilter)
	beego.InsertFilter("/*./user", beego.BeforeExec, filters.UserFilter)
	beego.InsertFilter("/*./user/:name", beego.BeforeExec, filters.UserFilter)
	beego.InsertFilter("/*./role", beego.BeforeExec, filters.UserRoleFilter)
	beego.InsertFilter("/*./role/:name", beego.BeforeExec, filters.UserRoleFilter)

	ns := beego.NewNamespace("/v1.0.0",
		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleContoller{},
			)),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			)),
		beego.NSNamespace("/authentication",
			beego.NSInclude(
				&controllers.AuthenticationController{},
			)),
		beego.NSNamespace("/gitlab",
			beego.NSInclude(
				&controllers.GitlabController{},
			)),
		beego.NSNamespace("/device",
			beego.NSInclude(
				&controllers.DeviceContoller{},
			)),
		beego.NSNamespace("/juniper",
			beego.NSInclude(
				&controllers.JuniperController{},
			)),
		beego.NSNamespace("/cisco",
			beego.NSInclude(
				&controllers.CiscoController{},
			)),
		beego.NSNamespace("/task",
			beego.NSInclude(
				&controllers.TaskController{},
			)),
	)
	beego.AddNamespace(ns)

}
