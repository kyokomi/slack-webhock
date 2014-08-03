package main

import (
	"github.com/kyokomi/go-gitlab-client/gogitlab"
	"testing"
)

func TestGetProjectName(t *testing.T) {

	gitlab := gogitlab.NewGitlab(GetGitlabBaseUrl(), apiPath, GetGitlabToken())

	projectName, err := GetProjectName(gitlab, 67040)
	if err != nil {
		t.Error(err)
	}
	if projectName != "gitlab-cli-test" {
		t.Error("not found projectName ", projectName)
	}
}

func TestGetUserName(t *testing.T) {

	gitlab := gogitlab.NewGitlab(GetGitlabBaseUrl(), apiPath, GetGitlabToken())

	userName, err := GetUserName(gitlab, 27917)
	if err != nil {
		t.Error(err)
	}
	if userName != "kyokomi" {
		t.Error("not found userName ", userName)
	}
}
