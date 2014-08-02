package main

import "os"

func GetGitlabToken() string {
	return os.Getenv("GITLAB_TOKEN")
}

func GetSlackToken() string {
	return os.Getenv("SLACK_TOKEN")
}

func GetSlackChannel() string {
	return os.Getenv("SLACK_CHANNEL")
}


