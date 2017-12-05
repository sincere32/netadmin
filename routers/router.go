// @APIVersion 1.0.0
// @Title NetAdmin API
// @Description Netadmin has every tool to get any job done, so codename for the new netadmin APIs.
// @Contact zengqi0529@163.com
package routers

import (
	"gitee.com/pippozq/netadmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//
	//beego.Include(&controllers.AuthorityController{})

	ns := beego.NewNamespace("/v1.0.0",
		beego.NSNamespace("/authority",
			beego.NSInclude(
				&controllers.AuthorityController{},
			)),
	)
	beego.AddNamespace(ns)

}
