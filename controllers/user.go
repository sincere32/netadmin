package controllers

import (
	"encoding/json"
	"gitee.com/pippozq/netadmin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"gitee.com/pippozq/netadmin/models"
)

type AuthorityController struct {
	utils.BaseController
}

// @Title Get Authority List
// @Description Get Authority
// @Success 200 {object} models.AuthorityList
// @Failure 404 None
// @Failure 405 Not Allowed
// @router / [get]
func (a *AuthorityController) GetAuthorityList() {
	o := orm.NewOrm()
	o.Using("default")

	var authority models.Authority
	var authorities []models.Authority
	o.QueryTable(authority).All(&authorities)

	a.ReturnOrmJson(0, int64(len(authorities)), authorities)
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

			if count, insertErr := o.Insert(authority); insertErr == nil {
				a.ReturnOrmJson(0, count, authority)
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
			if count, updateErr := o.Update(&query); updateErr == nil {
				a.ReturnOrmJson(0, count, query)
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

type UsersController struct {
	beego.Controller
}

func (u *UsersController) Get() {

	u.Data["Website"] = "beego.me"
	u.Data["Email"] = "astaxie@gmail.com"
	u.ServeJSON()
}

func (u *UsersController) Post() {
	o := orm.NewOrm()
	o.Using("default")
	var model interface{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &model)
	//beego.Info(fmt.Sprintf("%s", model["test"]))

}
