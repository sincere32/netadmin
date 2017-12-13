package schedules

import (
	"github.com/astaxie/beego/toolbox"
	"gitee.com/pippozq/netadmin/models"
	"time"
	"github.com/astaxie/beego/orm"
)

func InitTask(){
	o := orm.NewOrm()
	o.Using("default")

	var task models.Task
	var tasks []models.Task
	if _, err := o.QueryTable(task).Filter("enabled",true).All(&tasks); err == nil {
		for _, t := range tasks{
			s := new(Schedule)
			s.T = &t
			s.AddTask()
			s.RunTask()
		}
	}
}


type Schedule struct {
	Sch *toolbox.Task
	T    *models.Task
}


func (s *Schedule)AddTask(){
	if s.T.Type == "juniper"{
		if s.T.Module == "config"{
			s.Sch = toolbox.NewTask(s.T.Name, s.T.CronTime, s.JuniperConfigTask)
		}
	}
}

func (s *Schedule)RunTask(){
	nowTime := time.Now()
	s.Sch.SetNext(nowTime)
	toolbox.AddTask(s.T.Name,s.Sch)
	toolbox.StartTask()
}

func (s *Schedule)DeleteTask(){
	toolbox.DeleteTask(s.T.Name)
}

