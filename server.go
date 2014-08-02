package main

import (
	"github.com/kyokomi/go-gitlab-client/gogitlab"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"log"
	"net/http"
	"fmt"
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

	var text string

	// Project Get
	projectName, err := GetProjectName(gitlab, issues.ObjectAttributes.ProjectID)
	if err != nil {
		l.Println(err)
		return "Error"
	}
	text = text + fmt.Sprintln("プロジェクト名: ", projectName)

	// User Get Author
	authorName, err := GetUserName(gitlab, issues.ObjectAttributes.AuthorID)
	if err != nil {
		l.Println(err)
		return "Error"
	}
	text = text + fmt.Sprintln("作成者: ", authorName)

	// User Get Assignee
	if issues.ObjectAttributes.AssigneeID != 0 {
		assigneeName, err := GetUserName(gitlab, issues.ObjectAttributes.AssigneeID)
		if err != nil {
			l.Println(err)
			return "Error"
		}
		text = text + fmt.Sprintln("担当者: ", assigneeName)
	}

	text = text + fmt.Sprintln("タイトル: ", issues.ObjectAttributes.Title)
	text = text + fmt.Sprintln("内容: ", issues.ObjectAttributes.Description)
	text = text + fmt.Sprintln("作成日: ", issues.ObjectAttributes.CreatedAt)
	text = text + fmt.Sprintln("ステータス: ", issues.ObjectAttributes.State)
	l.Println(text)

	res, err := http.Get("https://slack.com/api/chat.postMessage?token=xoxp-2311903184-2311903186-2507218985-be003d&channel=C0295SK5L&text=" + issues.ObjectAttributes.Title + "&pretty=1")
	if err != nil {
		l.Println(err)
		return "Error"
	}
	l.Println(res)

	return "OK"
}

