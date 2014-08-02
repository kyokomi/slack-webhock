package main

import (
	"fmt"
	"net/http"
)

func SendMessage(message string) error {
	url := fmt.Sprintf("https://slack.com/api/chat.postMessage?token=%s&channel=%s&text=%s&pretty=1", GetSlackToken(), GetSlackChannel(), message)
	fmt.Println(url)
	_, err := http.Get(url)
	if err != nil {
		return err
	}
	return nil
}
