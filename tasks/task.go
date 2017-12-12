package tasks

import (
	"github.com/astaxie/beego/toolbox"
	"time"
	"gitee.com/pippozq/netadmin/models"
	"github.com/astaxie/beego"
)


func StartTask(t *toolbox.Task, name string){
	nowTime := time.Now()
	t.SetNext(nowTime)
	toolbox.AddTask(name, t)
	toolbox.StartTask()
	beego.Info(t.GetNext())
}



func Schedule(task *models.Task)  {
	t := new(toolbox.Task)
	//t.SetCron()
	if task.Type == "juniper"{
		beego.Info(task.Name,task.CronTime)
		t.SetCron(task.CronTime)

		//t := toolbox.NewTask(task.Name, task.CronTime, func() error {
		//	return JuniperConfigTask(task)
		//} )
		//StartTask(t, t.Taskname)
	}


}

