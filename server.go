package main

import (
	"github.com/kyokomi/go-gitlab-client/gogitlab"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"log"
)

const (
	baseUrl = "https://gitlab.com"
	apiPath = "/api/v3"
)

func main() {

	gitlab := gogitlab.NewGitlab(baseUrl, apiPath, GetGitlabToken())

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
	projectName, err := GetProjectName(gitlab, issues.ObjectAttributes.ProjectID)
	if err != nil {
		l.Println(err)
		return "Error"
	}

	// User Get Author
	authorName, err := GetUserName(gitlab, issues.ObjectAttributes.AuthorID)
	if err != nil {
		l.Println(err)
		return "Error"
	}

	// User Get Assignee
	var assigneeName string
	if issues.ObjectAttributes.AssigneeID != 0 {
		assigneeName, err = GetUserName(gitlab, issues.ObjectAttributes.AssigneeID)
		if err != nil {
			l.Println(err)
			return "Error"
		}
	}

	message := PostMessage{
		ProjectName  : projectName,
		AuthorName   : authorName,
		AssigneeName : assigneeName,
		Title        : issues.ObjectAttributes.Title,
		Description  : issues.ObjectAttributes.Description,
		CreatedAt    : issues.ObjectAttributes.CreatedAt,
		State        : issues.ObjectAttributes.State,
	}
	l.Println(message)

	if err := SendMessage(message); err != nil {
		l.Println(err)
		return "Error"
	}
	return "OK"
}

