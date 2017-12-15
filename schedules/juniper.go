package schedules

import (
	"encoding/json"
	"fmt"
	"github.com/pippozq/netadmin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"time"
)

func (s *Schedule) JuniperConfigTask() error {

	o := orm.NewOrm()
	o.Using("default")

	gitlab := models.Gitlab{ReposName: s.T.ReposName}

	history := new(models.TaskHistory)
	history.Task = s.T
	history.LastRunTime = time.Now()
	history.NextRunTime = s.Sch.GetNext()

	if o.Read(&gitlab, "ReposName") == orm.ErrNoRows {
		history.TaskMsg = "No Such Repository"
		history.LastSuccess = false

	} else {
		reposId, _ := gitlab.GetReposID(s.T.ReposName, gitlab.Url, gitlab.Token)
		blobId, _ := gitlab.GetFileBlobID(reposId, s.T.Branch, s.T.FilePath, gitlab.Url, gitlab.Token)
		fileContent, _ := gitlab.GetFileContent(reposId, blobId, gitlab.Url, gitlab.Token)

		juniperConfig := models.JuniperConfigPost{
			Hosts:       []string{s.T.Ip},
			User:        models.JuniperUser{Name: s.T.UserName, Password: s.T.Password},
			FileContent: fileContent,
		}

		req := httplib.Post(fmt.Sprintf("%s/juniper/config", beego.AppConfig.String("netadmin_driver_url")))
		req.Header("Content-Type", "application/json; charset=UTF-8")
		req.JSONBody(juniperConfig)
		response, err := req.Response()
		if err != nil {
			history.TaskMsg = err.Error()
			history.LastSuccess = false
		} else {
			if response.StatusCode == 200 {
				resultByte, _ := req.Bytes()
				err = json.Unmarshal(resultByte, &s.Message)
				if err == nil {
					var msg string
					for _, line := range s.Message["msg"].([]interface{}) {
						msg += line.(string)
					}
					history.LastSuccess = true
					history.TaskMsg = msg
				} else {
					history.LastSuccess = false
					history.TaskMsg = err.Error()
				}

			} else {
				history.LastSuccess = false
				history.TaskMsg = response.Status
			}
		}
	}

	if _, insertErr := o.Insert(history); insertErr == nil {
		return nil
	} else {
		return insertErr
	}

	return nil
}

func (s *Schedule) JuniperCommandTask() error {

	o := orm.NewOrm()
	o.Using("default")

	gitlab := models.Gitlab{ReposName: s.T.ReposName}

	history := new(models.TaskHistory)
	history.Task = s.T
	history.LastRunTime = time.Now()
	history.NextRunTime = s.Sch.GetNext()

	if o.Read(&gitlab, "ReposName") == orm.ErrNoRows {
		history.TaskMsg = "No Such Repository"
		history.LastSuccess = false

	} else {
		reposId, _ := gitlab.GetReposID(s.T.ReposName, gitlab.Url, gitlab.Token)
		blobId, _ := gitlab.GetFileBlobID(reposId, s.T.Branch, s.T.FilePath, gitlab.Url, gitlab.Token)
		fileContent, _ := gitlab.GetFileContent(reposId, blobId, gitlab.Url, gitlab.Token)

		juniperCommand := models.JuniperCommandRec{
			Hosts:   []string{s.T.Ip},
			User:    models.JuniperUser{Name: s.T.UserName, Password: s.T.Password},
			Command: fileContent,
		}

		req := httplib.Post(fmt.Sprintf("%s/juniper/command", beego.AppConfig.String("netadmin_driver_url")))
		req.Header("Content-Type", "application/json; charset=UTF-8")
		req.JSONBody(juniperCommand)
		response, err := req.Response()
		if err != nil {
			history.TaskMsg = err.Error()
			history.LastSuccess = false
		} else {
			if response.StatusCode == 200 {
				resultByte, _ := req.Bytes()
				err = json.Unmarshal(resultByte, &s.Message)
				if err == nil {
					var msg string
					for _, line := range s.Message["msg"].([]interface{}) {
						msg += line.(string)
					}
					history.LastSuccess = true
					history.TaskMsg = msg
				} else {
					history.LastSuccess = false
					history.TaskMsg = err.Error()
				}

			} else {
				history.LastSuccess = false
				history.TaskMsg = response.Status
			}
		}
	}

	if _, insertErr := o.Insert(history); insertErr == nil {
		return nil
	} else {
		return insertErr
	}

	return nil
}
