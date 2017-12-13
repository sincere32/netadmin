package schedules

import (
	"fmt"
	"gitee.com/pippozq/netadmin/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"time"
)

func InitTask() {
	o := orm.NewOrm()
	o.Using("default")

	var task models.Task
	var tasks []models.Task
	if _, err := o.QueryTable(task).Filter("enabled", true).All(&tasks); err == nil {
		for _, t := range tasks {
			s := new(Schedule)
			s.T = &t
			s.AddTask()
			s.RunTask()
		}
	}
}

type Message map[string]interface{}

type Schedule struct {
	Sch *toolbox.Task
	T   *models.Task
	Message
}

func (s *Schedule) AddTask() {
	if s.T.Type == "juniper" {
		if s.T.Module == "config" {
			s.Sch = toolbox.NewTask(fmt.Sprintf("%s_%s", s.T.Name, s.T.Ip), s.T.CronTime, s.JuniperConfigTask)
		} else if s.T.Module == "query" {
			s.Sch = toolbox.NewTask(fmt.Sprintf("%s_%s", s.T.Name, s.T.Ip), s.T.CronTime, s.JuniperCommandTask)
		}
	} else if s.T.Type == "cisco" {
		if s.T.Module == "config" {
			s.Sch = toolbox.NewTask(fmt.Sprintf("%s_%s", s.T.Name, s.T.Ip), s.T.CronTime, s.CiscoConfigTask)
		} else if s.T.Module == "query" {
			s.Sch = toolbox.NewTask(fmt.Sprintf("%s_%s", s.T.Name, s.T.Ip), s.T.CronTime, s.CiscoCommandTask)
		}
	}
}

func (s *Schedule) RunTask() {
	nowTime := time.Now()
	s.Sch.SetNext(nowTime)
	toolbox.AddTask(fmt.Sprintf("%s_%s", s.T.Name, s.T.Ip), s.Sch)
	toolbox.StartTask()
}

func (s *Schedule) DeleteTask() {
	toolbox.DeleteTask(fmt.Sprintf("%s_%s", s.T.Name, s.T.Ip))
}
