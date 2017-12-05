package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gitee.com/pippozq/netadmin/controllers:AuthorityController"] = append(beego.GlobalControllerRouter["gitee.com/pippozq/netadmin/controllers:AuthorityController"],
		beego.ControllerComments{
			Method: "GetAuthorityList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitee.com/pippozq/netadmin/controllers:AuthorityController"] = append(beego.GlobalControllerRouter["gitee.com/pippozq/netadmin/controllers:AuthorityController"],
		beego.ControllerComments{
			Method: "AddAuthority",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitee.com/pippozq/netadmin/controllers:AuthorityController"] = append(beego.GlobalControllerRouter["gitee.com/pippozq/netadmin/controllers:AuthorityController"],
		beego.ControllerComments{
			Method: "UpdateAuthority",
			Router: `/`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitee.com/pippozq/netadmin/controllers:AuthorityController"] = append(beego.GlobalControllerRouter["gitee.com/pippozq/netadmin/controllers:AuthorityController"],
		beego.ControllerComments{
			Method: "GetAuthority",
			Router: `/:name`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitee.com/pippozq/netadmin/controllers:AuthorityController"] = append(beego.GlobalControllerRouter["gitee.com/pippozq/netadmin/controllers:AuthorityController"],
		beego.ControllerComments{
			Method: "DeleteAuthority",
			Router: `/:name`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
