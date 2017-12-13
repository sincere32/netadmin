package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(Task), new(TaskHistory))
}

type TaskUser struct {
}

type Task struct {
	Id            int            `json:"id"`
	Created       time.Time      `orm:"auto_now_add;type(datetime)"`
	Name          string         `orm:"size(50);"json:"task_name"`
	ReposName     string         `orm:"size(100)"json:"repos_name"`
	Branch        string         `orm:"size(100)"json:"branch"`
	Ip            string         `orm:"size(100)"json:"device_ip"`
	CronTime      string         `orm:"size(100)"json:"cron_time"`
	Enabled       bool           `orm:"default(false)"json:"enabled"`
	UserName      string         `orm:"size(100)"json:"username"`
	Password      string         `orm:"size(255)"json:"password"`
	FilePath      string         `orm:"size(255)"json:"file_path"`
	Type          string         `orm:"size(20)"json:"type"`
	Module        string         `orm:"size(20)"json:"module"`
	TaskHistories []*TaskHistory `orm:"reverse(many)"json:"-"`
	LastRunTime   time.Time      `orm:"-"json:"last_run_time"`
	LastSuccess   bool           `orm:"-"json:"last_success"`
}

type TaskHistory struct {
	Id          int       `json:"id"`
	LastRunTime time.Time `orm:"null;default(null);type(datetime)"json:"last_run_time"`
	NextRunTime time.Time `orm:"null;default(null);type(datetime)"json:"next_run_time"`
	LastSuccess bool      `orm:"default(False)"json:"last_success"`
	TaskMsg     string    `orm:"size(255);null;default(null)"json:"msg"`
	Task        *Task     `orm:"rel(fk)"`
}
