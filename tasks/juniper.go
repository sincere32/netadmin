package tasks

import (
	"gitee.com/pippozq/netadmin/models"
	"github.com/astaxie/beego"
)

func JuniperConfigTask(task *models.Task) error{

	beego.Info(task.Name)
	return nil
}