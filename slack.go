package main

import (
	"fmt"
	"net/http"
	"net/url"
	"text/template"
	"bytes"
)

type TemplateExec struct {
	Message PostMessage
}

type PostMessage struct {
	ProjectName  string
	AuthorName   string
	AssigneeName string
	Title        string
	Descriptions []string
	CreatedAt    string
	State        string
	Milestone    string
}

const messageTemplate = `
_{{.Message.Title}}_ *{{.Message.State}}* {{.Message.CreatedAt}}

{{range $idx, $text := .Message.Descriptions}}
> {{$text}}{{end}}
ProjectName: *{{.Message.ProjectName}}* {{with .Message.Milestone}}({{.}}){{end}}
Author: <@{{.Message.AuthorName}}> Assignee: <@{{.Message.AssigneeName}}>
`

func SendMessage(message PostMessage) error {

	t := template.Must(template.New("message").Parse(messageTemplate))
	var buf bytes.Buffer
	if err := t.Execute(&buf, TemplateExec{Message: message}); err != nil {
		return err
	}

	postUrl := "https://slack.com/api/chat.postMessage"
	client := &http.Client{}
	res, err := client.PostForm(postUrl, url.Values{
		"token":      {GetSlackToken()},
		"channel":    {GetSlackChannel()},
		"text":       {buf.String()},
		"username":   {"Gitlab"},
		"icon_emoji": {":beers:"},
	})
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
