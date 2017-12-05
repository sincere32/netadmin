package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User), new(Authority))
}

type User struct {
	Id        int
	Name      string     `orm:"size(50);unique"json:"name"`
	Password  string     `orm:"size(255)"json:"password"`
	Email     string     `orm:"size(50);null"json:"email"`
	Tel       int        `orm:"null"json:"tel"`
	Authority *Authority `orm:"rel(fk);"json:"authority_id"`
}

type Authority struct {
	Id   int
	Name string `orm:"size(50);"json:"name"`
	Full bool   `orm:"default(False)"json:"full"`
	Rw   bool   `orm:"default(False)"json:"rw"`
}
