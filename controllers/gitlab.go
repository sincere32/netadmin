package controllers

import (
	"encoding/json"
	"fmt"
	"gitee.com/pippozq/netadmin/utils"
	"gitee.com/pippozq/netadmin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type GitlabController struct {
	utils.BaseController
}

// @Title Get Gitlab Commits
// @Description Add Role,if full and rw are false, then role is readonly
// @Param   post_body body  string  true "{ 'repos_name': 'project_name', 'file_name': 'file_name', 'file_content': 'file_content','branch':'branch','commit_message':'commit_message'}"
// @Success 200 {object} utils.AddGitlabCommitResp
// @Failure 404 body is empty
// @router /commits [post]
func (c *GitlabController) AddGitlabRepositoryCommits() {

	commits := new(models.AddGitlabCommitResp)

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &commits)

	if err != nil {
		c.ReturnJson(-1, err.Error())
	} else {

		gitlab := new(utils.Gitlab)
		reposId, _ := gitlab.GetReposID(commits.ReposName)

		action := new(models.GitlabCommitAddAction)
		action.Action = "create"
		action.Content = commits.FileContent
		action.FilePath = commits.FileName

		addCommits := new(models.GitlabAddCommit)

		addCommits.Actions = append(addCommits.Actions, *action)
		addCommits.Branch = commits.Branch
		addCommits.CommitMessage = commits.CommitMessage

		req := httplib.Post(fmt.Sprintf("%s/api/v4/projects/%d/repository/commits", beego.AppConfig.String("gitlab_url"), reposId))
		req.Header("Private-Token", beego.AppConfig.String("gitlab_token"))
		req.Header("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		req.JSONBody(addCommits)
		response, err := req.Response()

		if response.StatusCode == 200 || response.StatusCode == 201 {
			resp, _ := req.String()
			c.ReturnJson(0, resp)
		} else {
			beego.Error(-1, err.Error())
		}
	}

}
