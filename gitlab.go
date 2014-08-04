package main

import (
	"github.com/kyokomi/go-gitlab-client/gogitlab"
	"strconv"
)


func GetProjectName(gitlab *gogitlab.Gitlab, projectId int64) (string, error) {
	project, err := gitlab.Project(strconv.FormatInt(projectId, 10))
	if err != nil {
		return "", err
	}
	return project.Name, nil
}

func GetUserName(gitlab *gogitlab.Gitlab, userId int64) (string, error) {
	user, err := gitlab.User(strconv.FormatInt(userId, 10))
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

//	/projects/:id/milestones/:milestone_id
func GetMilestoneTitle(gitlab *gogitlab.Gitlab, projectId, milestoneId int64) (string, error) {
	milestone, err := gitlab.ProjectMilestone(strconv.FormatInt(projectId, 10), strconv.FormatInt(milestoneId, 10))
	if err != nil {
		return "", err
	}
	return milestone.Title, nil
}
