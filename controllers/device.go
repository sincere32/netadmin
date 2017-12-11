package controllers

import (
	"encoding/json"
	"gitee.com/pippozq/netadmin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"gitee.com/pippozq/netadmin/models"
)

type DeviceContoller struct {
	utils.BaseController
}

// @Title Get Device List
// @Description Get Device
// @Success 200 {object} []models.device
// @Failure 404 Not Found
// @router / [get]
func (c *DeviceContoller) GetDeviceList() {
	o := orm.NewOrm()
	o.Using("default")

	var dev models.Device
	var devices []models.Device
	if _, err := o.QueryTable(dev).All(&devices); err == nil {
		c.ReturnTableJson(0, len(devices), len(devices), devices)
	} else {
		c.ReturnJson(1, err.Error())
	}

}

// @Title Get Device Types
// @Description Get Device Type
// @Success 200 {object} models.Device
// @Failure 404 body is empty
// @router /types [get]
func (c *DeviceContoller) GetDeviceTypes() {
	o := orm.NewOrm()
	o.Using("default")

	var dev models.Device
	var devices []models.Device
	if _, err := o.QueryTable(dev).All(&devices); err == nil {
		var deviceType []string
		for _,d := range devices{
			has := false
			for _, dt := range deviceType{
				if d.Type == dt {
					has = true
					break
				}
			}
			if !has{
				deviceType = append(deviceType, d.Type)
			}

		}
		c.ReturnJson(0, deviceType)
	} else {
		c.ReturnJson(1, err.Error())
	}

}

// @Title Get Device By Type
// @Description Get Device By Type  "types" is key word!!!
// @Success 200 {object} models.Device
// @Param   type path  string  true  "device type"
// @Failure 404 body is empty
// @router /:type [get]
func (c *DeviceContoller) GetDevice() {
	o := orm.NewOrm()
	o.Using("default")

	deviceType := c.GetString(":type")

	var dev models.Device
	var devices []models.Device
	if _, err := o.QueryTable(dev).Filter("type",deviceType).All(&devices); err == nil {
		c.ReturnTableJson(0, len(devices), len(devices), devices)
	} else {
		c.ReturnJson(1, err.Error())
	}

}

// @Title Add Device
// @Description Add Device
// @Param   post_body body  string  true "{"name":'your name ',"super": false, "rw":false}"
// @Success 200 {object} models.Device
// @Failure 404 body is empty
// @router / [post]
func (c *DeviceContoller) AddDevice() {
	o := orm.NewOrm()
	o.Using("default")
	device := new(models.Device)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &device)
	beego.Info(device)
	if err != nil {
		c.ReturnJson(200, err.Error())
	} else {
		if o.Read(device, "Name") == orm.ErrNoRows && o.Read(device,"Ip") == orm.ErrNoRows{

			if _, insertErr := o.Insert(device); insertErr == nil {
				c.ReturnOrmJson(0, 1, device)
			} else {
				c.ReturnJson(1, insertErr.Error())
			}
		} else {
			c.ReturnJson(2, "This Name or Ip Already Exist")
		}

	}
}

// @Title Update Device
// @Description Update Device,if full and rw are false, then device is readonly
// @Param   patch_body body  string  true  "{"name":'your name ',"full": false, "rw":false}"
// @Success 200 {object} models.Device
// @Failure 404 body is empty
// @router / [patch]
func (c *DeviceContoller) UpdateDevice() {
	o := orm.NewOrm()
	o.Using("default")
	device := new(models.Device)
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &device)
	beego.Info(device)
	if err != nil {
		c.ReturnJson(200, err.Error())
	} else {
		query := models.Device{Ip: device.Ip}
		if o.Read(&query, "Ip") != orm.ErrNoRows {
			query.Name = device.Name
			query.Ip = device.Ip
			query.Type = device.Type
			if _, updateErr := o.Update(&query); updateErr == nil {
				c.ReturnOrmJson(0, 1, query)
			} else {
				c.ReturnJson(1, updateErr.Error())
			}
		} else {
			c.ReturnJson(2, "No Such Name Exist")
		}

	}
}

// @Title Delete Device
// @Description Delete Device
// @Param   name path  string  true  "device name"
// @Success 200 {object} models.Device
// @Failure 404 body is empty
// @router /:name [delete]
func (c *DeviceContoller) DeleteDevice() {
	o := orm.NewOrm()
	o.Using("default")

	name := c.GetString(":name")

	device := models.Device{Name: name}

	if o.Read(&device, "Name") != orm.ErrNoRows {
		if count, delErr := o.Delete(&device); delErr == nil {
			c.ReturnOrmJson(0, count, device)
		} else {
			c.ReturnOrmJson(1, 0, delErr.Error())
		}
	} else {
		c.ReturnJson(2, "No Such Name Exist")
	}
}
