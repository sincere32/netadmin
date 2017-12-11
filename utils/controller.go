package utils

import (
	"github.com/astaxie/beego"
)

type ReturnTable struct {
	Status          int         `json:"status"`
	RecordsTotal    int         `json:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered"`
	TableData       interface{} `json:"data"`
}

type ReturnJson struct {
	Status int         `json:"status"`
	Msg    interface{} `json:"msg"`
}

type Message map[string]interface{}

type Messages []Message

type BaseController struct {
	beego.Controller
	Message
	Messages
}

// Format Return Body

func (base *BaseController) ReturnTableJson(status, total, records int, data interface{}) {

	base.Data["json"] = ReturnTable{
		Status:          status,
		RecordsTotal:    total,
		RecordsFiltered: records,
		TableData:       data,
	}
	base.ServeJSON()
}

func (base *BaseController) ReturnNetDriverJson(msg interface{}) {
	base.Data["json"] = msg
	base.ServeJSON()
}

func (base *BaseController) ReturnJson(status int, msg interface{}) {
	base.Data["json"] = ReturnJson{
		Status: status,
		Msg:    msg,
	}
	base.ServeJSON()
}

func (base *BaseController) ReturnOrmJson(status int, count int64, msg interface{}) {
	type ReturnOrmJ struct {
		Status int         `json:"status"`
		Count  int64       `json:"count"`
		Msg    interface{} `json:"msg"`
	}

	base.Data["json"] = ReturnOrmJ{
		Status: status,
		Count:  count,
		Msg:    msg,
	}
	base.ServeJSON()
}

func (base *BaseController) Error404() {
	base.Data["json"] = ReturnJson{
		Status: 404,
		Msg:    "Not Found",
	}
	base.ServeJSON()
}

func (base *BaseController) Error500() {
	base.Data["json"] = ReturnJson{
		Status: 500,
		Msg:    "Internal Error",
	}
	base.ServeJSON()
}
