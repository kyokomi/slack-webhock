package main

import (
	"fmt"
	"net/http"
	"net/url"
	"text/template"
	"bytes"
)

type PostMessage struct {
	ProjectName  string
	AuthorName   string
	AssigneeName string
	Title        string
	Description  string
	CreatedAt    string
	State        string
}

const messageTemplate = `
_プロジェクト名: {{.ProjectName}}_
{{.Title}} {{.Description}} {{.CreatedAt}} *{{.State}}*
> 作成者: <@{{.AuthorName}}>
> 担当者: <@{{.AssigneeName}}>
`

func SendMessage(message PostMessage) error {

	t := template.Must(template.New("message").Parse(messageTemplate))
	var buf bytes.Buffer
	if err := t.Execute(&buf, message); err != nil {
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
