package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Gitlab))
}

type GitlabProject struct {
	Id        int    `json:"id"`
	ReposName string `json:"name"`
	ReposPath string `json:"path"`
}

type GitlabFile struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
}

type AddGitlabCommitResp struct {
	ReposName     string `json:"repos_name"`
	FileName      string `json:"file_name"`
	FileContent   string `json:"file_content"`
	Branch        string `json:"branch"`
	CommitMessage string `json:"commit_message"`
}

type GitlabCommitAddUpdateAction struct {
	Action   string `json:"action"`
	FilePath string `json:"file_path"`
	Content  string `json:"content"`
}

type GitlabCommitAddOrUpdate struct {
	Actions       []GitlabCommitAddUpdateAction `json:"actions"`
	Branch        string                        `json:"branch"`
	CommitMessage string                        `json:"commit_message"`
}

type GitlabCommitDel struct {
	Actions       []GitlabCommitDelAction `json:"actions"`
	Branch        string                  `json:"branch"`
	CommitMessage string                  `json:"commit_message"`
}

type GitlabCommitDelAction struct {
	Action   string `json:"action"`
	FilePath string `json:"file_path"`
}

type Gitlab struct {
	Id        int    `json:"id"`
	Url       string `orm:"size(255)"json:"url"`
	User      string `orm:"size(20);null"json:"user"`
	Token     string `orm:"size(255)"json:"token"`
	ReposName string `orm:"size(100)"json:"repos_name"`
}

type GitlabInterface interface {
	GetReposID(reposName string) (id int, err error)
	GetFileBlobID(reposID int, ref, filePath string) (blobID string, err error)
	GetFileContent(reposID int, blobId string) (content string, err error)
}

func (g *Gitlab) GetReposID(reposName string) (id int, err error) {
	req := httplib.Get(fmt.Sprintf("%s/api/v4/projects", beego.AppConfig.String("gitlab_url")))
	req.Header("Private-Token", beego.AppConfig.String("gitlab_token"))
	req.Header("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	response, err := req.Response()
	beego.Debug(response)
	var projects []GitlabProject
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
	var files []GitlabFile
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
