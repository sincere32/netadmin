package utils

import (
	"encoding/json"
	"fmt"
	"gitee.com/pippozq/netadmin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type Gitlab struct {
	models.Gitlab
}

func (g *Gitlab) GetReposID(reposName string) (id int, err error) {
	req := httplib.Get(fmt.Sprintf("%s/api/v4/projects", beego.AppConfig.String("gitlab_url")))
	req.Header("Private-Token", beego.AppConfig.String("gitlab_token"))
	req.Header("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	response, err := req.Response()
	beego.Debug(response)
	var projects []models.GitlabProject
	if response.StatusCode == 200 {
		resp, _ := req.Bytes()
		err := json.Unmarshal(resp, &projects)
		if err != nil {
			beego.Error(err)
			return id, err
		} else {
			for _, p := range projects {
				beego.Info(p, p.ReposName, reposName)
				if p.ReposName == reposName {
					return p.Id, err
				}
			}
		}
	} else {
		beego.Error(err)
	}
	return id, err
}

func (g *Gitlab) GetFileContent(reposID int, blobId string) (content string, err error) {
	beego.Info(fmt.Sprintf("%s/api/v4/projects/%d/repository/blobs/%s/raw", beego.AppConfig.String("gitlab_url"), reposID, blobId))
	req := httplib.Get(fmt.Sprintf("%s/api/v4/projects/%d/repository/blobs/%s/raw", beego.AppConfig.String("gitlab_url"), reposID, blobId))
	req.Header("Private-Token", beego.AppConfig.String("gitlab_token"))
	req.Header("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	response, err := req.Response()
	beego.Debug(response)

	if response.StatusCode == 200 {
		content, _ := req.String()
		if err != nil {
			beego.Error(err)
			return content, err
		} else {
			return content, err
		}
	} else {
		beego.Error(err)
	}
	return content, err
}

func (g *Gitlab) GetFileBlobID(reposID int, ref, filePath string) (blobID string, err error) {
	beego.Info(fmt.Sprintf("%s/api/v4/projects/%d/repository/tree?recursive=true&ref=%s", beego.AppConfig.String("gitlab_url"), reposID, ref))
	req := httplib.Get(fmt.Sprintf("%s/api/v4/projects/%d/repository/tree?recursive=true&ref=%s", beego.AppConfig.String("gitlab_url"), reposID, ref))
	req.Header("Private-Token", beego.AppConfig.String("gitlab_token"))
	req.Header("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	response, err := req.Response()
	beego.Debug(response)
	var files []models.GitlabFile
	if response.StatusCode == 200 {
		resp, _ := req.Bytes()
		err := json.Unmarshal(resp, &files)
		if err != nil {
			beego.Error(err)
			return blobID, err
		} else {
			for _, f := range files {

				if f.Path == filePath {
					return f.Id, err
				}
			}
		}
	} else {
		beego.Error(err)
	}
	return blobID, err
}
