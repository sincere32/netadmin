package utils

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}


type FormatJsonInterface interface {
	EncodeJson(status int, msg interface{})
}

// Format Return Body
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