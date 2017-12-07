package utils

import (
	"github.com/astaxie/beego"
)

type ReturnTable struct {
	RecordsTotal    int         `json:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered"`
	TableData       interface{} `json:"data"`
}

type BaseController struct {
	beego.Controller
}

type FormatJsonInterface interface {
	EncodeJson(status int, msg interface{})
}

// Format Return Body

func (base *BaseController) ReturnTableJson(total, records int, data interface{}) {

	base.Data["json"] = ReturnTable{
		RecordsTotal:    total,
		RecordsFiltered: records,
		TableData:       data,
	}
	base.ServeJSON()
}

func (base *BaseController) ReturnJson(status int, msg interface{}) {
	type ReturnJson struct {
		Status int         `json:"status"`
		Msg    interface{} `json:"msg"`
	}

	base.Data["json"] = ReturnJson{
		Status: status,
		Msg:    msg,
	}
	base.ServeJSON()
}

func (base *BaseController) ReturnOrmJson(status int, count int64, msg interface{}) {
	type ReturnJson struct {
		Status int         `json:"status"`
		Count  int64       `json:"count"`
		Msg    interface{} `json:"msg"`
	}

	base.Data["json"] = ReturnJson{
		Status: status,
		Count:  count,
		Msg:    msg,
	}
	base.ServeJSON()
}

func (base *BaseController) Error404() {
	type ReturnJson struct {
		Status int         `json:"status"`
		Msg    interface{} `json:"msg"`
	}
	base.Data["json"] = ReturnJson{
		Status: 404,
		Msg:    "Not Found",
	}
	base.ServeJSON()
}

func (base *BaseController) Error500() {
	type ReturnJson struct {
		Status int         `json:"status"`
		Msg    interface{} `json:"msg"`
	}
	base.Data["json"] = ReturnJson{
		Status: 500,
		Msg:    "Internal Error",
	}
	base.ServeJSON()
}
