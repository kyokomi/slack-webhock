package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"log"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Post("/issues-events", binding.Bind(IssuesEvents{}), doIssuesEvents)
	m.Run()
}

func doIssuesEvents(l *log.Logger, issues IssuesEvents) string {

	l.Println(issues)

	return "OK"
}
