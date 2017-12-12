package controllers

import (
	"encoding/json"
	"fmt"
	"gitee.com/pippozq/netadmin/models"
	"gitee.com/pippozq/netadmin/utils"
	"gitee.com/pippozq/netadmin/tasks"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	if _, err := o.QueryTable(task).RelatedSel().All(&tasks); err == nil {
		c.ReturnTableJson(0, len(tasks), len(tasks), tasks)
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
				o.QueryTable(task).Filter("name",task.Name).Filter("ip",task.Ip).All(&query)
				if len(query) == 0{
					tasks.Schedule(task)
					if _, insertErr := o.Insert(task); insertErr == nil {
						addList = append(addList, task)
						beego.Info(*task)
					} else {
						addErr = true
						errList = append(errList, task)
					}
				}else {
					beego.Info(fmt.Sprintf("This Name Already Exist:%s", task.Name))
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
// @router / [post]
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
				tasks.Schedule(task)
				o.QueryTable(task).Filter("name",task.Name).Filter("ip",task.Ip).All(&query)
				if len(query) == 1{
					if _, updateErr := o.Update(task); updateErr == nil {
						addList = append(addList, task)
					} else {
						addErr = true
						errList = append(errList, task)
					}
				}else {
					beego.Info(fmt.Sprintf("This Name Not Exist:%s", task.Name))
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

	//if err != nil {
	//	c.ReturnJson(-1, err.Error())
	//} else {
	//
	//}
	//tk1 := toolbox.NewTask("tk1", "0/10 * * * * *", test)
	//nowTime := time.Now()
	//beego.Info(nowTime)
	//tk1.SetNext(nowTime)
	//toolbox.AddTask("tk1", tk1)
	//toolbox.StartTask()

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
	beego.Info(name, ip)
	task := models.Task{Name: name,Ip:ip}

	if o.Read(&task, "Name","Ip") != orm.ErrNoRows {
		if count, delErr := o.Delete(&task); delErr == nil {
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
	if _, err := o.QueryTable(th).Filter("task_id",id).All(&ths); err == nil {
		c.ReturnTableJson(0, len(ths), len(ths), ths)
	} else {
		c.ReturnJson(1, err.Error())
	}

}

func (c *TaskController)CheckTimeStr(){
	if err := recover(); err != nil {
		c.ReturnJson(-1, err)
	}
}