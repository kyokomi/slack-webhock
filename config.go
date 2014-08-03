package main

import "os"

func GetGitlabBaseUrl() string {
	return os.Getenv("GITLAB_BASE_URL")
}

func GetGitlabToken() string {
	return os.Getenv("GITLAB_TOKEN")
}

func GetSlackToken() string {
	return os.Getenv("SLACK_TOKEN")
}

func GetSlackChannel() string {
	return os.Getenv("SLACK_CHANNEL")
}
