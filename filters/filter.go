package filters

import (
	"github.com/pippozq/netadmin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"strings"
)

func ReturnJson(ctx *context.Context, status int, msg interface{}) {
	type ReturnJson struct {
		Status int         `json:"status"`
		Msg    interface{} `json:"msg"`
	}
	ctx.Output.JSON(ReturnJson{status, msg}, true, true)
}

func LoginFilter(ctx *context.Context) {
	o := orm.NewOrm()
	o.Using("default")
	url := strings.Split(ctx.Request.RequestURI, "/")

	if url[2] != "authentication" {
		name := ctx.GetCookie("nickname")
		if name == "" {
			ReturnJson(ctx, 1, "No Authentication Expired")
		} else {
			user := models.User{Name: name}
			if o.Read(&user, "Name") != orm.ErrNoRows {
				ses := ctx.Input.CruSession.Get(beego.AppConfig.String("login_session"))
				cookie := ctx.GetCookie(beego.AppConfig.String("login_session"))
				if ses != nil {
					if ses != cookie {
						ReturnJson(ctx, -1, "Authentication Error")
					}
				} else {
					ReturnJson(ctx, 1, "No Authentication Expired")
				}
			} else {
				ReturnJson(ctx, 2, "No Such User Exist")
			}
		}

	}

}

func RoleFilter(ctx *context.Context) {
	o := orm.NewOrm()
	o.Using("default")
	url := strings.Split(ctx.Request.RequestURI, "/")
	if url[2] != "authentication" {
		if ctx.Request.Method == "POST" || ctx.Request.Method == "PATCH" || ctx.Request.Method == "DELETE" {
			name := ctx.GetCookie("nickname")
			var user models.User
			var query []models.User

			o.QueryTable(user).Filter("name", name).RelatedSel().All(&query)
			if len(query) > 0 {
				super := query[0].Role.Super
				rw := query[0].Role.Rw

				if !super && !rw {
					ReturnJson(ctx, 1, "Authority is ReadOnly")
				}
			} else {
				ReturnJson(ctx, -1, "No Such User")
			}

		}

	}

}

func UserFilter(ctx *context.Context) {
	o := orm.NewOrm()
	o.Using("default")

	name := ctx.GetCookie("nickname")
	var user models.User
	var query []models.User

	o.QueryTable(user).Filter("name", name).RelatedSel().All(&query)

	super := query[0].Role.Super

	if !super {
		ReturnJson(ctx, 1, "Authority is Not Super")
	}
}

func UserRoleFilter(ctx *context.Context) {
	o := orm.NewOrm()
	o.Using("default")

	name := ctx.GetCookie("nickname")
	var user models.User
	var query []models.User

	o.QueryTable(user).Filter("name", name).RelatedSel().All(&query)

	super := query[0].Role.Super
	if !super {
		ReturnJson(ctx, 1, "Authority is Not Super")
	}
}
