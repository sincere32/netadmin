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
