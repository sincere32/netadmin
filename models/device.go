package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Device))
}

type Device struct {
	Id   int    `json:"id"`
	Name string `orm:"size(50);unique"json:"device_name"`
	Type string `orm:"size(50)"json:"device_type"`
	Ip   string `orm:"size(255);null"json:"ip"`
}
