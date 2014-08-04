package main

import (
	"testing"
	"strings"
)

func TestSendMessage(t *testing.T) {


	message := PostMessage{
		ProjectName  : "Slack-Webhock-Gitlab Project",
		AuthorName   : "kyokomi",
		AssigneeName : "kyokomi",
		Title        : "Demo Issue Title",
		Descriptions : strings.Split("hogehogehoge.\nfugafugafuga?\npiyopiyopiyo!!", "\n"),
		CreatedAt    : "2014/08/03 7:53:00",
		State        : "opend",
		Milestone    : "v0.0.1",
	}
	err := SendMessage(message)

	if err != nil {
		t.Error(err)
	}
}
