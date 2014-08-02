package main

import "testing"

func TestSendMessage(t *testing.T) {
	err := SendMessage("testOK!")
	if err != nil {
		t.Error(err)
	}
}
