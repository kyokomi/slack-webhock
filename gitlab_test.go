package main

import (
	"testing"
	"github.com/kyokomi/go-gitlab-client/gogitlab"
)


func TestGetProjectName(t *testing.T) {

	gitlab := gogitlab.NewGitlab(baseUrl, apiPath, token)

	projectName, err := GetProjectName(gitlab, 67040)
	if err != nil {
		t.Error(err)
	}
	if projectName != "gitlab-cli-test" {
		t.Error("not found projectName ", projectName)
	}
}

func TestGetUserName(t *testing.T) {

	gitlab := gogitlab.NewGitlab(baseUrl, apiPath, token)

	userName, err := GetUserName(gitlab, 27917)
	if err != nil {
		t.Error(err)
	}
	if userName != "kyokomi" {
		t.Error("not found userName ", userName)
	}
}
