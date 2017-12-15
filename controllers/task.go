package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/pippozq/netadmin/models"
	"github.com/pippozq/netadmin/schedules"
	"github.com/pippozq/netadmin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type TaskController struct {
	utils.BaseController
}

// @Title Get Task List
// @Description Get Task List
// @Success 200 {object} []models.task
// @Failure 404 Not Found
// @router / [get]
func (c *TaskController) GetTaskList() {
	o := orm.NewOrm()
	o.Using("default")

	var task models.Task

	var tasks []models.Task

	var taskHistory models.TaskHistory
	var taskHistories []*models.TaskHistory

	var returnTable []models.Task

	if _, err := o.QueryTable(task).All(&tasks); err == nil {

		for _, t := range tasks {
			o.QueryTable(taskHistory).Filter("Task", t).RelatedSel().OrderBy("-id").All(&taskHistories)
			if len(taskHistories) == 0 {
				t.LastSuccess = false
				t.LastRunTime = time.Now()
			} else {
				t.LastSuccess = taskHistories[0].LastSuccess
				t.LastRunTime = taskHistories[0].LastRunTime
			}
			returnTable = append(returnTable, t)
		}

		c.ReturnTableJson(0, len(returnTable), len(returnTable), returnTable)
	} else {
		c.ReturnJson(1, err.Error())
	}

}

// @Title Add Task
// @Description Add Task
// @Param   post_body body  string  true "{"name":'your name ',"super": false, "rw":false}"
// @Success 200 {object} models.Task
// @Failure 404 Not Found
// @router / [post]
func (c *TaskController) AddTaskController() {
	o := orm.NewOrm()
	o.Using("default")
	var ts map[string][]string
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ts)
	if err != nil {
		c.ReturnJson(-1, err.Error())
	} else {
		addErr := false
		var addList, errList []*models.Task
		defer c.CheckTimeStr()
		for _, t := range ts["schedules"] {

			task := new(models.Task)
			err := json.Unmarshal([]byte(t), task)
			if err != nil {
				c.ReturnJson(-1, err.Error())
			} else {

				var query []models.Task
				o.QueryTable(task).Filter("name", task.Name).Filter("ip", task.Ip).All(&query)
				if len(query) == 0 {
					s := new(schedules.Schedule)
					s.T = task
					s.AddTask()
					if _, insertErr := o.Insert(task); insertErr == nil {
						addList = append(addList, task)
						s.RunTask()
					} else {
						addErr = true
						errList = append(errList, task)
					}
				} else {
					beego.Error(fmt.Sprintf("This Name Already Exist:%s", task.Name))
					errList = append(errList, task)
				}
			}

		}
		if addErr {
			c.ReturnJson(-1, errList)
		} else {
			c.ReturnJson(0, addList)
		}

	}

}

// @Title Update Task
// @Description Update Task
// @Param   post_body body  string  true "{"name":'your name ',"super": false, "rw":false}"
// @Success 200 {object} models.Task
// @Failure 404 Not Found
// @router / [patch]
func (c *TaskController) UpdateTaskController() {
	o := orm.NewOrm()
	o.Using("default")
	var ts map[string][]string
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ts)
	if err != nil {
		c.ReturnJson(-1, err.Error())
	} else {
		defer c.CheckTimeStr()
		addErr := false
		var addList, errList []*models.Task
		for _, t := range ts["schedules"] {
			task := new(models.Task)
			err := json.Unmarshal([]byte(t), task)
			if err != nil {
				c.ReturnJson(-1, err.Error())
			} else {
				var query []models.Task
				o.QueryTable(task).Filter("name", task.Name).Filter("ip", task.Ip).All(&query)
				if len(query) == 1 {

					query[0].UserName = task.UserName
					query[0].Password = task.Password
					query[0].CronTime = task.CronTime
					query[0].Enabled = task.Enabled

					s := new(schedules.Schedule)
					s.T = &query[0]
					s.AddTask()

					if _, updateErr := o.Update(&query[0]); updateErr == nil {
						addList = append(addList, task)

						// enable or disable
						if query[0].Enabled {
							s.RunTask()
						} else {
							s.DeleteTask()
						}

					} else {
						addErr = true
						beego.Error(updateErr)
						errList = append(errList, task)
					}
				} else {
					beego.Error(fmt.Sprintf("This Name Not Exist:%s", task.Name))
					errList = append(errList, task)
				}
			}

		}
		if addErr {
			c.ReturnJson(-1, errList)
		} else {
			c.ReturnJson(0, addList)
		}

	}
}

// @Title Delete Task
// @Description Delete Device
// @Param   name query  string  true  "task name"
// @Param   ip query  string  true  "device ip"
// @Success 200 {object} models.Device
// @Failure 404 body is empty
// @router / [delete]
func (c *TaskController) DeleteTaskController() {
	o := orm.NewOrm()
	o.Using("default")

	name := c.GetString("name")
	ip := c.GetString("ip")
	task := models.Task{Name: name, Ip: ip}

	if o.Read(&task, "Name", "Ip") != orm.ErrNoRows {
		if count, delErr := o.Delete(&task); delErr == nil {
			s := new(schedules.Schedule)
			s.T = &task
			s.DeleteTask()
			c.ReturnOrmJson(0, count, task)
		} else {
			c.ReturnOrmJson(1, 0, delErr.Error())
		}
	} else {
		c.ReturnJson(2, "No Such Name Exist")
	}
}

// @Title Get Task History  List
// @Description Get Task List
// @Param id path int true "task id"
// @Success 200 {object} []models.TaskHistory
// @Failure 404 Not Found
// @router /:id/history [get]
func (c *TaskController) GetTaskHistory() {
	o := orm.NewOrm()
	o.Using("default")

	id := c.GetString(":id")

	var th models.TaskHistory
	var ths []models.TaskHistory
	if _, err := o.QueryTable(th).Filter("task_id", id).OrderBy("-last_run_time").All(&ths); err == nil {
		c.ReturnTableJson(0, len(ths), len(ths), ths)
	} else {
		c.ReturnJson(1, err.Error())
	}

}

func (c *TaskController) CheckTimeStr() {
	if err := recover(); err != nil {
		c.ReturnJson(-1, err)
	}
}
