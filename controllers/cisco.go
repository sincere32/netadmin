package controllers

import (
	"gitee.com/pippozq/netadmin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"

	"encoding/json"
	"fmt"
	"gitee.com/pippozq/netadmin/models"
	"github.com/astaxie/beego/orm"
)

type CiscoController struct {
	utils.BaseController
}

// @Title Cisco Command
// @Description Run command
// @Param   post_body body  string  true "{"hosts":'[host1,host2,host3]',"user": {'name':'name','password':'password'},"command": command}"
// @Success 200 {object} models.CiscoCommandRec
// @Failure 404 body is emptys
// @router /command [post]
func (c *CiscoController) Command() {

	req := httplib.Post(fmt.Sprintf("%s/cisco/command", beego.AppConfig.String("netadmin_driver_url")))
	req.Header("Content-Type", "application/json; charset=UTF-8")
	req.Body(c.Ctx.Input.RequestBody)
	response, err := req.Response()
	if err != nil {
		c.ReturnJson(-1, err.Error())
	} else {
		if response.StatusCode == 200 {
			resultByte, _ := req.Bytes()
			err = json.Unmarshal(resultByte, &c.Message)
			if err == nil {
				c.ReturnNetDriverJson(c.Message)
			} else {
				beego.Error(err)
				c.ReturnJson(-1, err.Error())
			}

		}
	}
}

// @Title  Cisco Config
// @Description Execute Cisco Script
// @Param   post_body body  string  true "{"hosts":'[host1,host2,host3]',"user": {'name':'name','password':'password'},'repos_name': 'repos_name','ref':'ref','file_path':'file_path}"
// @Success 200 {object} models.CiscoConfigRec
// @Failure 404 body is empty
// @router /config [post]
func (c *CiscoController) Config() {

	rec := new(models.CiscoConfigRec)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &rec)

	if err == nil {

		o := orm.NewOrm()
		o.Using("default")

		gitlab := models.Gitlab{ReposName: rec.ReposName}

		if o.Read(&gitlab, "ReposName") == orm.ErrNoRows {
			c.ReturnJson(2, "No such repository")
		} else {
			reposId, _ := gitlab.GetReposID(rec.ReposName, gitlab.Url, gitlab.Token)
			blobId, _ := gitlab.GetFileBlobID(reposId, rec.Ref, rec.FilePath, gitlab.Url, gitlab.Token)
			fileContent, _ := gitlab.GetFileContent(reposId, blobId, gitlab.Url, gitlab.Token)

			ciscoConfig := models.CiscoConfigPost{
				Hosts:       rec.Hosts,
				User:        rec.User,
				FileContent: fileContent,
				BlobId:blobId,
			}

			req := httplib.Post(fmt.Sprintf("%s/cisco/config", beego.AppConfig.String("netadmin_driver_url")))
			req.Header("Content-Type", "application/json; charset=UTF-8")
			req.JSONBody(ciscoConfig)
			response, err := req.Response()
			if err != nil {
				c.ReturnJson(-1, err.Error())
			} else {
				if response.StatusCode == 200 {
					resultByte, _ := req.Bytes()
					err = json.Unmarshal(resultByte, &c.Message)
					if err == nil {
						c.ReturnNetDriverJson(c.Message)
					} else {
						c.ReturnJson(-1, err.Error())
					}

				} else {
					c.ReturnJson(1, response.Status)
				}
			}
		}

	}else{
		c.ReturnJson(-1, err.Error())
	}

}
