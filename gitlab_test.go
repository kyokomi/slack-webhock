package main

import (
	"testing"
	"github.com/kyokomi/go-gitlab-client/gogitlab"
	"strconv"
)

func TestGetProjectName(t *testing.T) {

	gitlab := gogitlab.NewGitlab(GetGitlabBaseUrl(), apiPath, GetGitlabToken())

	projects, err := gitlab.Projects()
	if err != nil {
		t.Error(err)
	}
	if len(projects) == 0 {
		return
	}

	var project *gogitlab.Project
	project = projects[0]

	if _, err := GetProjectName(gitlab, project.Id); err != nil {
		t.Error(err)
	}
}

func TestGetUserName(t *testing.T) {

	gitlab := gogitlab.NewGitlab(GetGitlabBaseUrl(), apiPath, GetGitlabToken())

	users, err := gitlab.Users()
	if err != nil {
		t.Error(err)
	}
	if len(users) == 0 {
		return
	}

	var user *gogitlab.User
	user = users[0]

	if _, err := GetUserName(gitlab, user.Id); err != nil {
		t.Error(err)
	}
}

func TestGetMilestoneTitle(t *testing.T) {

	gitlab := gogitlab.NewGitlab(GetGitlabBaseUrl(), apiPath, GetGitlabToken())

	projects, err := gitlab.Projects()
	if err != nil {
		t.Error(err)
	}
	if len(projects) == 0 {
		return
	}

	var project *gogitlab.Project
	project = projects[0]

	milestones, err := gitlab.ProjectMilestones(strconv.Itoa(project.Id))
	if err != nil {
		t.Error(err)
	}
	if len(milestones) == 0 {
		return
	}

	var milestone *gogitlab.Milestone
	milestone = milestones[0]

	if _, err := GetMilestoneTitle(gitlab, project.Id, milestone.ID); err != nil {
		t.Error(err)
	}
}
