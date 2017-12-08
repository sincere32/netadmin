package controllers

import (
	"encoding/json"
	"gitee.com/pippozq/netadmin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"gitee.com/pippozq/netadmin/models"

	"time"
)

type RoleContoller struct {
	utils.BaseController
}

// @Title Get Role List
// @Description Get Role
// @Success 200 {object} []models.Role
// @Failure 404 Not Found
// @router / [get]
func (a *RoleContoller) GetRoleList() {
	o := orm.NewOrm()
	o.Using("default")

	var role models.Role
	var authorities []models.Role
	o.QueryTable(role).All(&authorities)

	a.ReturnTableJson(len(authorities), len(authorities), authorities)
}

// @Title Get Role
// @Description Get Role
// @Success 200 {object} models.Role
// @Param   name path  string  true  "role name"
// @Failure 404 body is empty
// @router /:name [get]
func (a *RoleContoller) GetRole() {
	o := orm.NewOrm()
	o.Using("default")

	name := a.GetString(":name")

	role := models.Role{Name: name}

	if o.Read(&role, "Name") == orm.ErrNoRows {
		a.ReturnJson(1, "none")
	} else {
		a.ReturnOrmJson(0, 1, role)
	}

}

// @Title Add Role
// @Description Add Role,if full and rw are false, then role is readonly
// @Param   post_body body  string  true "{"name":'your name ',"super": false, "rw":false}"
// @Success 200 {object} models.Role
// @Failure 404 body is empty
// @router / [post]
func (a *RoleContoller) AddRole() {
	o := orm.NewOrm()
	o.Using("default")
	role := new(models.Role)
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &role)
	beego.Info(role)
	if err != nil {
		a.ReturnJson(200, err.Error())
	} else {
		if o.Read(role, "Name") == orm.ErrNoRows {

			if _, insertErr := o.Insert(role); insertErr == nil {
				a.ReturnOrmJson(0, 1, role)
			} else {
				a.ReturnJson(1, insertErr.Error())
			}
		} else {
			a.ReturnJson(2, "This Name Already Exist")
		}

	}
}

// @Title Update Role
// @Description Update Role,if full and rw are false, then role is readonly
// @Param   patch_body body  string  true  "{"name":'your name ',"full": false, "rw":false}"
// @Success 200 {object} models.Role
// @Failure 404 body is empty
// @router / [patch]
func (a *RoleContoller) UpdateRole() {
	o := orm.NewOrm()
	o.Using("default")
	role := new(models.Role)
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &role)
	beego.Info(role)
	if err != nil {
		a.ReturnJson(200, err.Error())
	} else {
		query := models.Role{Name: role.Name}
		if o.Read(&query, "Name") != orm.ErrNoRows {
			query.Rw = role.Rw
			query.Super = role.Super
			if _, updateErr := o.Update(&query); updateErr == nil {
				a.ReturnOrmJson(0, 1, query)
			} else {
				a.ReturnJson(1, updateErr.Error())
			}
		} else {
			a.ReturnJson(2, "No Such Name Exist")
		}

	}
}

// @Title Delete Role
// @Description Delete Role
// @Param   name path  string  true  "role name"
// @Success 200 {object} models.Role
// @Failure 404 body is empty
// @router /:name [delete]
func (a *RoleContoller) DeleteRole() {
	o := orm.NewOrm()
	o.Using("default")

	name := a.GetString(":name")

	role := models.Role{Name: name}

	if o.Read(&role, "Name") != orm.ErrNoRows {
		if count, delErr := o.Delete(&role); delErr == nil {
			a.ReturnOrmJson(0, count, role)
		} else {
			a.ReturnOrmJson(1, 0, delErr.Error())
		}
	} else {
		a.ReturnJson(2, "No Such Name Exist")
	}
}

type UserController struct {
	utils.BaseController
}

// @Title Get User List
// @Description Get User
// @Success 200 {object} []utils.ReturnTable
// @Failure 404 None
// @Failure 405 Not Allowed
// @router / [get]
func (a *UserController) GetUserList() {
	o := orm.NewOrm()
	o.Using("default")

	var user models.User
	var users []models.User
	o.QueryTable(user).RelatedSel().All(&users)

	a.ReturnTableJson(len(users), len(users), users)
}

// @Title Get User
// @Description Get User
// @Success 200 {object} models.User
// @Param   name path  string  true  "User name"
// @Failure 404 body is empty
// @router /:name [get]
func (a *UserController) GetUser() {
	o := orm.NewOrm()
	o.Using("default")

	name := a.GetString(":name")

	var user models.User
	var users []models.User

	o.QueryTable(user).Filter("name", name).RelatedSel().All(&users)
	a.ReturnOrmJson(0, int64(len(users)), users)

}

// @Title Add User
// @Description Add User,name password role_id is required
// @Param   post_body body  string  true "{"name":'your name ',"password": 'password', "role_id":'id',"tel":'12345678901', "email":'xx@ff.com'}"
// @Success 200 {object} models.User
// @Failure 404 body is empty
// @router / [post]
func (a *UserController) AddUser() {
	o := orm.NewOrm()
	o.Using("default")

	user := new(models.User)
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &user)
	beego.Info(user)
	if err != nil {
		a.ReturnJson(200, err.Error())
	} else {
		if o.Read(user, "Name") == orm.ErrNoRows {
			if user.Role == nil {
				a.ReturnJson(-1, "No Role")
			} else {
				role := new(models.Role)
				role = user.Role

				if o.Read(role) == orm.ErrNoRows {
					a.ReturnJson(-1, "No such role")
				} else {

					if _, insertErr := o.Insert(user); insertErr == nil {
						a.ReturnOrmJson(0, 1, user)
					} else {
						a.ReturnJson(1, insertErr.Error())
					}
				}

			}
		} else {
			a.ReturnJson(2, "This Name Already Exist")
		}

	}
}

// @Title Update User
// @Description Update User,User,name password role_id is required
// @Param   patch_body body  string  true  "{"name":'your name ',"password": 'password', "role_id":id,"tel":12345678901, "email":'xx@ff.com'}"
// @Success 200 {object} models.User
// @Failure 404 body is empty
// @router / [patch]
func (a *UserController) UpdateUser() {
	o := orm.NewOrm()
	o.Using("default")
	user := new(models.User)
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &user)
	beego.Info(user)
	if err != nil {
		a.ReturnJson(200, err.Error())
	} else {
		query := models.User{Name: user.Name}
		if o.Read(&query, "Name") != orm.ErrNoRows {
			query.Name = user.Name
			query.Password = user.Password
			query.Email = user.Email
			query.Tel = user.Tel

			role := new(models.Role)
			role = user.Role

			if o.Read(role) == orm.ErrNoRows {
				a.ReturnJson(-1, "No such role")
			} else {
				query.Role = role
				if _, updateErr := o.Update(&query); updateErr == nil {
					a.ReturnOrmJson(0, 1, query)
				} else {
					a.ReturnJson(1, updateErr.Error())
				}
			}
		} else {
			a.ReturnJson(2, "No Such Name Exist")
		}

	}
}

// @Title Delete User
// @Description Delete User
// @Param   name path  string  true  "User name"
// @Success 200 {object} models.User
// @Failure 404 body is empty
// @router /:name [delete]
func (a *UserController) DeleteUser() {
	o := orm.NewOrm()
	o.Using("default")

	name := a.GetString(":name")

	user := models.User{Name: name}

	if o.Read(&user, "Name") != orm.ErrNoRows {
		if count, delErr := o.Delete(&user); delErr == nil {
			a.ReturnOrmJson(0, count, user)
		} else {
			a.ReturnOrmJson(1, 0, delErr.Error())
		}
	} else {
		a.ReturnJson(2, "No Such Name Exist")
	}
}

type AuthenticationController struct {
	utils.BaseController
}

// @Title SignIn
// @Description SignIn
// @Param   post_body body  string  true "{"name":'your name ',"password": 'password'}"
// @Success 200  {"status": "-1 error 0 succeed 1 failed", msg:""}
// @Failure 404 body is empty
// @router / [post]
func (a *AuthenticationController) SignIn() {
	o := orm.NewOrm()
	o.Using("default")

	user := new(models.User)
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &user)
	if err == nil {
		var query []models.User

		qs := o.QueryTable(user)

		// Check User Exist

		qs.Filter("name", user.Name).All(&query)

		if len(query) == 0 {
			a.ReturnJson(2, "No such User")
		} else {
			if query[0].Password != user.Password {
				a.ReturnJson(3, "Password is Wrong")
			} else {
				sess := a.StartSession()
				defer sess.SessionRelease(a.Ctx.ResponseWriter)
				s := a.CruSession.Get(beego.AppConfig.String("login_session"))
				if s == nil {
					nowTime := time.Now()
					sValue := nowTime.Format("20170102150405")
					a.CruSession.Set(beego.AppConfig.String("login_session"), sValue)
					a.Ctx.SetCookie(beego.AppConfig.String("login_session"), sValue)
					a.Ctx.SetCookie("nickname", query[0].Name)
					a.ReturnJson(0, "Authentication Success")
				} else {
					a.ReturnJson(1, "Expired")
				}

			}

		}
	} else {
		a.ReturnJson(-1, err.Error())
	}
}

// @Title SignOut
// @Description SignOut
// @Param   name path  string  true "me"
// @Success 200 {"status":" 0 success 1 error", "msg":""}
// @Failure 404 body is empty
// @router /:name [delete]
func (a *AuthenticationController) SignOut() {
	o := orm.NewOrm()
	o.Using("default")

	name := a.GetString(":name")

	var user models.User
	var query []models.User

	o.QueryTable(user).Filter("name", name).RelatedSel().All(&query)

	if len(query) == 0 {
		a.ReturnJson(2, "No such User")
	} else {
		sess := a.StartSession()
		defer sess.SessionRelease(a.Ctx.ResponseWriter)
		ses := a.CruSession.Get(beego.AppConfig.String("login_session"))
		if ses != nil {
			a.DestroySession()
			a.CruSession.Flush()
			a.ReturnJson(0, "Authentication Delete")
		} else {
			a.ReturnJson(1, "No Authentication Expired")
		}
	}

}

// @Title Check Online Status
// @Description Check Online Status
// @Param  name path  string  true "me"
// @Success 200 {"status":"0 expired 1 not expired 2 no user", "msg":""}
// @Failure 404 not found
// @router /:name [get]
func (a *AuthenticationController) CheckLogin() {

	o := orm.NewOrm()
	o.Using("default")

	name := a.GetString(":name")

	user := models.User{Name: name}

	if o.Read(&user, "Name") != orm.ErrNoRows {
		sess := a.StartSession()
		defer sess.SessionRelease(a.Ctx.ResponseWriter)
		ses := a.CruSession.Get(beego.AppConfig.String("login_session"))
		cookie := a.Ctx.GetCookie(beego.AppConfig.String("login_session"))
		if ses != nil {
			if ses == cookie {
				a.ReturnJson(0, "Authentication Expired")
			} else {
				a.ReturnJson(-1, "Authentication Error")
			}

		} else {
			a.ReturnJson(1, "No Authentication Expired")
		}
	} else {
		a.ReturnJson(2, "No Such User Exist")
	}

}
