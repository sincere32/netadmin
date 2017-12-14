package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User), new(Role))
}

type User struct {
	Id       int    `json:"id"`
	Name     string `orm:"size(50);unique"json:"name"`
	Password string `orm:"size(255)"json:"password"`
	Email    string `orm:"size(50);null"json:"email"`
	Tel      string `orm:"size(11);null"json:"tel"`
	Role     *Role  `orm:"rel(fk);"json:"role"`
}

type Role struct {
	Id    int    `json:"id"`
	Name  string `orm:"size(50);"json:"name"`
	Super bool   `orm:"default(False)"json:"super"`
	Rw    bool   `orm:"default(False)"json:"rw"`
}

func InitUser() bool {
	o := orm.NewOrm()
	o.Using("default")

	role := Role{Name: "admin", Super: true}
	if o.Read(&role, "Name") == orm.ErrNoRows {
		if _, err := o.Insert(&role); err != nil {
			return false
		}
	} else {
		if _, err := o.Update(&role); err != nil {
			return false
		}
	}

	admin := User{Name: "admin", Password: "696d29e0940a4957748fe3fc9efd22a3", Role: &role}
	if o.Read(&admin, "Name") == orm.ErrNoRows {
		if _, err := o.Insert(&admin); err != nil {
			return false
		}
	} else {
		admin.Password = "696d29e0940a4957748fe3fc9efd22a3"
		if _, err := o.Update(&admin); err != nil {
			return false
		}
	}
	return true
}
