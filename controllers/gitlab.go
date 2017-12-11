package controllers

import (
	"encoding/json"
	"fmt"
	"gitee.com/pippozq/netadmin/models"
	"gitee.com/pippozq/netadmin/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
)

type GitlabController struct {
	utils.BaseController
}

// @Title Add Gitlab Commits
// @Description Push Commits
// @Param   post_body body  string  true "{ 'repos_name': 'project_name', 'file_name': 'file_name', 'file_content': 'file_content','branch':'branch','commit_message':'commit_message'}"
// @Success 200 {object} models.AddGitlabCommitResp
// @Failure 404 {object} utils.ReturnJson
// @router /repository/commits [post]
func (c *GitlabController) AddGitlabRepositoryCommits() {

	commits := new(models.AddGitlabCommitResp)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &commits)
	if err != nil {
		c.ReturnJson(-1, err.Error())
	} else {

		gitlab := models.Gitlab{ReposName: commits.ReposName}

		o := orm.NewOrm()
		o.Using("default")

		if o.Read(&gitlab, "ReposName") == orm.ErrNoRows {
			c.ReturnJson(2, "No such repository")
		} else {
			beego.Info("fffff")

			reposId, _ := gitlab.GetReposID(commits.ReposName, gitlab.Url, gitlab.Token)
			action := new(models.GitlabCommitAddUpdateAction)
			action.Action = "create"
			action.Content = commits.FileContent
			action.FilePath = commits.FileName

			addCommits := new(models.GitlabCommitAddOrUpdate)

			addCommits.Actions = append(addCommits.Actions, *action)
			addCommits.Branch = commits.Branch
			addCommits.CommitMessage = commits.CommitMessage

			req := httplib.Post(fmt.Sprintf("%s/api/v4/projects/%d/repository/commits", gitlab.Url, reposId))
			req.Header("Private-Token", gitlab.Token)
			req.Header("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
			req.JSONBody(addCommits)
			beego.Info(addCommits)
			beego.Info(req.String())
			response, _ := req.Response()
			resp, err := req.Bytes()
			if err == nil {
				if je := json.Unmarshal(resp, &c.Message); je == nil {
					if response.StatusCode == 200 || response.StatusCode == 201 {
						c.ReturnJson(0, c.Message)
					} else {
						c.ReturnJson(1, c.Message)
					}
				}
			} else {
				beego.Error(-1, err.Error())
			}
		}
	}

}

// @Title Update Gitlab Commits
// @Description Add Role,if full and rw are false, then role is readonly
// @Param   post_body body  string  true "{ 'repos_name': 'project_name', 'file_name': 'file_name', 'file_content': 'file_content','branch':'branch','commit_message':'commit_message'}"
// @Success 200 {object} utils.AddGitlabCommitResp
// @Failure 404 {object} utils.ReturnJson
// @router /repository/commits [patch]
func (c *GitlabController) UpdateGitlabRepositoryCommits() {

	commits := new(models.AddGitlabCommitResp)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &commits)

	if err != nil {
		c.ReturnJson(-1, err.Error())
	} else {

		gitlab := models.Gitlab{ReposName: commits.ReposName}

		o := orm.NewOrm()
		o.Using("default")

		if o.Read(&gitlab, "ReposName") == orm.ErrNoRows {
			c.ReturnJson(2, "No such repository")
		} else {
			reposId, _ := gitlab.GetReposID(commits.ReposName, gitlab.Url, gitlab.Token)

			action := new(models.GitlabCommitAddUpdateAction)
			action.Action = "update"
			action.Content = commits.FileContent
			action.FilePath = commits.FileName

			addCommits := new(models.GitlabCommitAddOrUpdate)

			addCommits.Actions = append(addCommits.Actions, *action)
			addCommits.Branch = commits.Branch
			addCommits.CommitMessage = commits.CommitMessage

			req := httplib.Post(fmt.Sprintf("%s/api/v4/projects/%d/repository/commits", gitlab.Url, reposId))
			req.Header("Private-Token", gitlab.Token)
			req.Header("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
			req.JSONBody(addCommits)
			response, _ := req.Response()
			resp, err := req.Bytes()
			if err == nil {
				if je := json.Unmarshal(resp, &c.Message); je == nil {
					if response.StatusCode == 200 || response.StatusCode == 201 {
						c.ReturnJson(0, c.Message)
					} else {
						c.ReturnJson(1, c.Message)
					}
				}
			} else {
				beego.Error(-1, err.Error())
			}
		}

	}

}

// @Title Delete Gitlab Commits
// @Description Add Role,if full and rw are false, then role is readonly
// @Param   repos_name query string true
// @Param   branch query string true
// @Param   file_name query string true
// @Param   commit_message query string true
// @Success 200 {object} utils.ReturnJson
// @Failure 404 {object} utils.ReturnJson
// @router /repository/commits [delete]
func (c *GitlabController) DeleteGitlabRepositoryCommits() {

	reposName := c.GetString("repos_name")
	fileName := c.GetString("file_name")
	branch := c.GetString("branch", "master")
	commitMessage := c.GetString("commit_message", "delete")

	gitlab := models.Gitlab{ReposName: reposName}

	o := orm.NewOrm()
	o.Using("default")

	if o.Read(&gitlab, "ReposName") == orm.ErrNoRows {
		c.ReturnJson(2, "No such repository")
	} else {
		reposId, err := gitlab.GetReposID(reposName, gitlab.Url, gitlab.Token)

		action := new(models.GitlabCommitDelAction)
		action.Action = "delete"
		action.FilePath = fileName

		delCommits := new(models.GitlabCommitDel)

		delCommits.Actions = append(delCommits.Actions, *action)
		delCommits.Branch = branch
		delCommits.CommitMessage = commitMessage

		req := httplib.Post(fmt.Sprintf("%s/api/v4/projects/%d/repository/commits", gitlab.Url, reposId))
		req.Header("Private-Token", gitlab.Token)
		req.Header("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		req.JSONBody(delCommits)
		response, _ := req.Response()
		resp, err := req.Bytes()
		if err == nil {

			if je := json.Unmarshal(resp, &c.Message); je == nil {
				if response.StatusCode == 200 || response.StatusCode == 201 {
					c.ReturnJson(0, c.Message)
				} else {
					c.ReturnJson(1, c.Message)
				}
			}
		} else {
			beego.Error(-1, err.Error())
		}
	}

}

// @Title Get Gitlab File Blob
// @Description Add Role,if full and rw are false, then role is readonly
// @Param   repos_name query  string  true
// @Param   blob_id query  string  true
// @Success 200 {object} utils.ReturnJson
// @Failure 404 {object} utils.ReturnJson
// @router /repository/blobs [get]
func (c *GitlabController) GitlabRepositoryFileBlobs() {

	o := orm.NewOrm()
	o.Using("default")

	reposName := c.GetString("repos_name")
	blobId := c.GetString("blob_id")

	gitlab := models.Gitlab{ReposName: reposName}

	if o.Read(&gitlab, "ReposName") == orm.ErrNoRows {
		c.ReturnJson(2, "No such repository")
	} else {
		reposId, err := gitlab.GetReposID(reposName, gitlab.Url, gitlab.Token)
		req := httplib.Get(fmt.Sprintf("%s/api/v4/projects/%d/repository/blobs/%s/raw", gitlab.Url, reposId, blobId))
		req.Header("Private-Token", gitlab.Token)
		req.Header("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		response, _ := req.Response()
		resp, err := req.String()

		if err == nil {
			if response.StatusCode == 200 {
				c.ReturnJson(0, resp)
			} else {
				errResp, _ := req.Bytes()
				json.Unmarshal(errResp, &c.Message)
				c.ReturnJson(1, c.Message)
			}
		} else {
			beego.Error(-1, err.Error())
			c.ReturnJson(-1, err.Error())
		}

	}

}

// @Title Get Gitlab Tree
// @Description Add Role,if full and rw are false, then role is readonly
// @Param   repos_name query  string  true
// @Param   branch query  string  true
// @Success 200 {object} []utils.ReturnTable
// @Failure 404 {object} utils.ReturnJson
// @router /repository/tree [get]
func (c *GitlabController) GitlabRepositoryTree() {

	o := orm.NewOrm()
	o.Using("default")

	reposName := c.GetString("repos_name")
	branch := c.GetString("branch")

	gitlab := models.Gitlab{ReposName: reposName}

	if o.Read(&gitlab, "ReposName") == orm.ErrNoRows {
		c.ReturnJson(2, "No such repository")
	} else {
		reposId, err := gitlab.GetReposID(reposName, gitlab.Url, gitlab.Token)
		req := httplib.Get(fmt.Sprintf("%s/api/v4/projects/%d/repository/tree?ref=%s&recursive=true", gitlab.Url, reposId, branch))
		req.Header("Private-Token", gitlab.Token)
		req.Header("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		response, _ := req.Response()
		resp, err := req.Bytes()

		if err == nil {
			if response.StatusCode == 200 {
				json.Unmarshal(resp, &c.Messages)
				c.ReturnTableJson(0, len(c.Messages), len(c.Messages), c.Messages)
			} else {
				json.Unmarshal(resp, &c.Message)
				c.ReturnJson(1, c.Message)
			}
		} else {
			beego.Error(-1, err.Error())
		}

	}

}

// @Title Get Gitlab Branch
// @Description Add Role,if full and rw are false, then role is readonly
// @Param   repos_name query  string  true
// @Success 200 {object} []utils.ReturnTable
// @Failure 404 {object} utils.ReturnJson
// @router /repository/branches [get]
func (c *GitlabController) GitlabRepositoryBranch() {

	o := orm.NewOrm()
	o.Using("default")

	reposName := c.GetString("repos_name")

	gitlab := models.Gitlab{ReposName: reposName}

	if o.Read(&gitlab, "ReposName") == orm.ErrNoRows {
		c.ReturnJson(2, "No such repository")
	} else {
		reposId, err := gitlab.GetReposID(reposName, gitlab.Url, gitlab.Token)
		req := httplib.Get(fmt.Sprintf("%s/api/v4/projects/%d/repository/branches", gitlab.Url, reposId))
		req.Header("Private-Token", gitlab.Token)
		req.Header("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		response, _ := req.Response()
		resp, err := req.Bytes()
		if err == nil {
			if response.StatusCode == 200 {
				json.Unmarshal(resp, &c.Messages)
				c.ReturnJson(0, c.Messages)
			} else {
				json.Unmarshal(resp, &c.Message)
				c.ReturnJson(1, c.Message)
			}
		} else {
			beego.Error(-1, err.Error())
		}

	}

}

// @Title Get Gitlab Repository List
// @Description Get Gitlab Repository List
// @Success 200 {object} []models.Gitlab
// @Failure 404 {object} utils.ReturnJson
// @router /repository [get]
func (c *GitlabController) GetGitlabRepositoryList() {

	o := orm.NewOrm()
	o.Using("default")

	var repos models.Gitlab
	var query []models.Gitlab

	_, err := o.QueryTable(repos).All(&query)

	if err != nil {
		c.ReturnJson(-1, err.Error())
	} else {
		c.ReturnTableJson(0, len(query), len(query), query)
	}

}

// @Title Add Gitlab Repository
// @Description Add Gitlab Repository
// @Param   post_body body  string  true "{ 'repos_name': 'url', 'url': 'user', 'token': 'token'}"
// @Success 200 {object} models.Gitlab
// @Failure 404 {object} utils.ReturnJson
// @router /repository [post]
func (c *GitlabController) AddGitlabRepository() {
	o := orm.NewOrm()
	o.Using("default")

	gitlab := new(models.Gitlab)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &gitlab)

	if err != nil {
		c.ReturnJson(0, err.Error())
	} else {
		if o.Read(gitlab, "ReposName") == orm.ErrNoRows {
			if _, insertErr := o.Insert(gitlab); insertErr == nil {
				c.ReturnOrmJson(0, 1, gitlab)
			} else {
				c.ReturnJson(1, insertErr.Error())
			}
		} else {
			c.ReturnJson(2, "This Name Already Exist")
		}
	}

}
