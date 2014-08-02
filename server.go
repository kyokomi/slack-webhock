package main

import (
	"github.com/kyokomi/go-gitlab-client/gogitlab"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"log"
	"strconv"
)

const (
	baseUrl = "https://gitlab.com"
	apiPath = "/api/v3"
	// TODO: あとで環境変数にする
	token = "T5giqDC8dtz9ukx9D16B"
)
func main() {

	gitlab := gogitlab.NewGitlab(baseUrl, apiPath, token)

	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Map(gitlab)

	m.Post("/issues-events", binding.Bind(IssuesEvents{}), doIssuesEvents)
	m.Run()
}

func doIssuesEvents(l *log.Logger, issues IssuesEvents, gitlab *gogitlab.Gitlab) string {

	l.Println(issues)

	// Project Get
	project, err := gitlab.Project(strconv.FormatInt(issues.ObjectAttributes.ProjectID, 20))
	if err != nil {
		l.Println(err)
		return "Error"
	}
	l.Println("プロジェクト名: ", project.Name)

	// User Get Author
	author, err := gitlab.User(strconv.FormatInt(issues.ObjectAttributes.AuthorID, 20))
	if err != nil {
		l.Println(err)
		return "Error"
	}
	l.Println("作成者: ", author.Name)

	// User Get Assignee
	if issues.ObjectAttributes.AssigneeID != 0 {
		assignee, err := gitlab.User(strconv.FormatInt(issues.ObjectAttributes.AssigneeID, 20))
		if err != nil {
			l.Println(err)
			return "Error"
		}
		l.Println("担当者: ", assignee.Name)
	}

	l.Println("タイトル: ", issues.ObjectAttributes.Title)
	l.Println("内容: ", issues.ObjectAttributes.Description)
	l.Println("作成日: ", issues.ObjectAttributes.CreatedAt)
	l.Println("ステータス: ", issues.ObjectAttributes.State)

	return "OK"
}

