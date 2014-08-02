package main

import "testing"

func TestSendMessage(t *testing.T) {

	message := PostMessage{
		ProjectName  : "AAAAAA PJ",
		AuthorName   : "kyokomi",
		AssigneeName : "kyokomi",
		Title        : "Test",
		Description  : "hogehogefuga",
		CreatedAt    : "2014/08/03 7:53",
		State        : "opend",
	}
	err := SendMessage(message)

	if err != nil {
		t.Error(err)
	}
}
