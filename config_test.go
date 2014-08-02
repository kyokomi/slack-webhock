package main

import "testing"

func TestGetGitlabToken(t *testing.T) {
	token := GetGitlabToken()
	if token == "" {
		t.Error("get gitlab token error")
	}
}

func TestGetSlackToken(t *testing.T) {
	token := GetSlackToken()
	if token == "" {
		t.Error("get slack token error")
	}
}

func TestGetSlackChannel(t *testing.T) {
	token := GetSlackChannel()
	if token == "" {
		t.Error("get slack channel error")
	}
}


