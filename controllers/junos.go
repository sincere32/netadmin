package controllers

import (
	"gitee.com/pippozq/netadmin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"

	"encoding/json"
	"fmt"
)

type JunosController struct {
	utils.BaseController
}

type NetDriverResp struct {
	Status int         `json:"status"`
	Msg    interface{} `json:"msg"`
}

// @Title Add Role
// @Description Add Role,if full and rw are false, then role is readonly
// @Param   post_body body  string  true "{"name":'your name ',"super": false, "rw":false}"
// @Success 200 {object} models.JunosCommand
// @Failure 404 body is empty
// @router /command [post]
func (c *JunosController) Command() {

	url := fmt.Sprintf("%s/junos/command", beego.AppConfig.String("netadmin_driver_url"))
	beego.Info(url)
	req := httplib.Post(url)
	req.Header("Content-Type", "application/json; charset=UTF-8")
	req.Body(c.Ctx.Input.RequestBody)
	response, err := req.Response()
	if err != nil {
		c.ReturnJson(-1, err.Error())
	} else {
		if response.StatusCode == 200 {
			resultByte, _ := req.Bytes()
			var ndr NetDriverResp
			err = json.Unmarshal(resultByte, &ndr)
			if err == nil {
				c.ReturnNetDriverJson(ndr)
			} else {
				beego.Error(err)
			}

		}
	}
}

// @Title Add Role
// @Description Add Role,if full and rw are false, then role is readonly
// @Param   post_body body  string  true "{"name":'your name ',"super": false, "rw":false}"
// @Success 200 {object} models.JunosCommand
// @Failure 404 body is empty
// @router /config [post]
func (c *JunosController) Config() {

	url := fmt.Sprintf("%s/junos/config", beego.AppConfig.String("netadmin_driver_url"))
	beego.Info(url)
	req := httplib.Post(url)
	req.Header("Content-Type", "application/json; charset=UTF-8")
	req.Body(c.Ctx.Input.RequestBody)
	response, err := req.Response()
	if err != nil {
		c.ReturnJson(-1, err.Error())
	} else {
		if response.StatusCode == 200 {
			resultByte, _ := req.Bytes()
			var ndr NetDriverResp
			err = json.Unmarshal(resultByte, &ndr)
			if err == nil {
				c.ReturnNetDriverJson(ndr)
			} else {
				beego.Error(err)
			}

		}
	}
}
