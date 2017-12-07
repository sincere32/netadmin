package controllers

import (
	"encoding/json"
	"gitee.com/pippozq/netadmin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"gitee.com/pippozq/netadmin/models"

	"time"
)

type AuthorityController struct {
	utils.BaseController
}

// @Title Get Authority List
// @Description Get Authority
// @Success 200 {object} []models.Authority
// @Failure 404 Not Found
// @router / [get]
func (a *AuthorityController) GetAuthorityList() {
	o := orm.NewOrm()
	o.Using("default")

	var authority models.Authority
	var authorities []models.Authority
	o.QueryTable(authority).All(&authorities)

	a.ReturnTableJson(len(authorities), len(authorities), authorities)
}

// @Title Get Authority
// @Description Get Authority
// @Success 200 {object} models.Authority
// @Param   name path  string  true  "authority name"
// @Failure 404 body is empty
// @router /:name [get]
func (a *AuthorityController) GetAuthority() {
	o := orm.NewOrm()
	o.Using("default")

	name := a.GetString(":name")

	authority := models.Authority{Name: name}

	if o.Read(&authority, "Name") == orm.ErrNoRows {
		a.ReturnJson(1, "none")
	} else {
		a.ReturnOrmJson(0, 1, authority)
	}

}

// @Title Add Authority
// @Description Add Authority,if full and rw are false, then authority is readonly
// @Param   post_body body  string  true "{"name":'your name ',"full": false, "rw":false}"
// @Success 200 {object} models.Authority
// @Failure 404 body is empty
// @router / [post]
func (a *AuthorityController) AddAuthority() {
	o := orm.NewOrm()
	o.Using("default")
	authority := new(models.Authority)
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &authority)
	if err != nil {
		a.ReturnJson(200, err.Error())
	} else {
		if o.Read(authority, "Name") == orm.ErrNoRows {

			if _, insertErr := o.Insert(authority); insertErr == nil {
				a.ReturnOrmJson(0, 1, authority)
			} else {
				a.ReturnJson(0, insertErr.Error())
			}
		} else {
			a.ReturnJson(1, "This Name Already Exist")
		}

	}
}

// @Title Update Authority
// @Description Update Authority,if full and rw are false, then authority is readonly
// @Param   patch_body body  string  true  "{"name":'your name ',"full": false, "rw":false}"
// @Success 200 {object} models.Authority
// @Failure 404 body is empty
// @router / [patch]
func (a *AuthorityController) UpdateAuthority() {
	o := orm.NewOrm()
	o.Using("default")
	authority := new(models.Authority)
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &authority)
	beego.Info(authority)
	if err != nil {
		a.ReturnJson(200, err.Error())
	} else {
		query := models.Authority{Name: authority.Name}
		if o.Read(&query, "Name") != orm.ErrNoRows {
			query.Rw = authority.Rw
			query.Full = authority.Full
			if _, updateErr := o.Update(&query); updateErr == nil {
				a.ReturnOrmJson(0, 1, query)
			} else {
				a.ReturnJson(0, updateErr.Error())
			}
		} else {
			a.ReturnJson(1, "No Such Name Exist")
		}

	}
}

// @Title Delete Authority
// @Description Delete Authority
// @Param   name path  string  true  "authority name"
// @Success 200 {object} models.Authority
// @Failure 404 body is empty
// @router /:name [delete]
func (a *AuthorityController) DeleteAuthority() {
	o := orm.NewOrm()
	o.Using("default")

	name := a.GetString(":name")

	authority := models.Authority{Name: name}

	if o.Read(&authority, "Name") != orm.ErrNoRows {
		if count, delErr := o.Delete(&authority); delErr == nil {
			a.ReturnOrmJson(0, count, authority)
		} else {
			a.ReturnOrmJson(0, 0, delErr.Error())
		}
	} else {
		a.ReturnJson(1, "No Such Name Exist")
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
// @Description Add User,name password authority_id is required
// @Param   post_body body  string  true "{"name":'your name ',"password": 'password', "authority_id":'id',"tel":'12345678901', "email":'xx@ff.com'}"
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
			if user.Authority == nil {
				a.ReturnJson(-1, "No Authority")
			} else {
				authority := new(models.Authority)
				authority = user.Authority

				if o.Read(authority) == orm.ErrNoRows {
					a.ReturnJson(-1, "No such authority")
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
// @Description Update User,User,name password authority_id is required
// @Param   patch_body body  string  true  "{"name":'your name ',"password": 'password', "authority_id":id,"tel":12345678901, "email":'xx@ff.com'}"
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

			authority := new(models.Authority)
			authority = user.Authority

			if o.Read(authority) == orm.ErrNoRows {
				a.ReturnJson(-1, "No such authority")
			} else {
				query.Authority = authority
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

// @Title Login
// @Description Login
// @Param   post_body body  string  true "{"name":'your name ',"password": 'password'}"
// @Success 200  {"status": "-1 error 0 succeed 1 failed", msg:""}
// @Failure 404 body is empty
// @router / [post]
func (a *AuthenticationController) Login() {
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
			a.ReturnJson(1, "No such User")
		} else {
			if query[0].Password != user.Password {
				a.ReturnJson(1, "Password is Wrong")
			} else {
				sess := a.StartSession()
				defer sess.SessionRelease(a.Ctx.ResponseWriter)
				s := a.CruSession.Get(beego.AppConfig.String("login_session"))
				if s == nil {
					nowTime := time.Now()
					sValue := nowTime.Format("20170102150405")
					a.CruSession.Set(beego.AppConfig.String("login_session"), sValue)
					a.Ctx.SetCookie(beego.AppConfig.String("login_session"), sValue)
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

// @Title Logout
// @Description Logout
// @Param   name path  string  true "me"
// @Success 200 {"status":" 0 success 1 error", "msg":""}
// @Failure 404 body is empty
// @router /:name [delete]
func (a *AuthenticationController) Logout() {
	o := orm.NewOrm()
	o.Using("default")

	name := a.GetString(":name")

	var user models.User
	var query []models.User

	o.QueryTable(user).Filter("name", name).RelatedSel().All(&query)

	if len(query) == 0 {
		a.ReturnJson(1, "No such User")
	} else {
		sess := a.StartSession()
		defer sess.SessionRelease(a.Ctx.ResponseWriter)
		ses := a.CruSession.Get(beego.AppConfig.String("login_session"))
		if ses != nil {
			a.DestroySession()
			a.CruSession.Flush()
			a.ReturnJson(1, "Authentication Delete")
		} else {
			a.ReturnJson(0, "No Authentication Expired")
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
				a.ReturnJson(0, "Authentication Error")
			}

		} else {
			a.ReturnJson(1, "No Authentication Expired")
		}
	} else {
		a.ReturnJson(2, "No Such User Exist")
	}

}
