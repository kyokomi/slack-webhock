package main

type IssuesEvents struct {
	ObjectAttributes struct {
		AssigneeID  int64       `json:"assignee_id"`
		AuthorID    int64       `json:"author_id"`
		BranchName  interface{} `json:"branch_name"`
		CreatedAt   string      `json:"created_at"`
		Description string      `json:"description"`
		ID          int64       `json:"id"`
		Iid         int64       `json:"iid"`
		MilestoneID int64       `json:"milestone_id"`
		Position    int64       `json:"position"`
		ProjectID   int64       `json:"project_id"`
		State       string      `json:"state"`
		Title       string      `json:"title"`
		UpdatedAt   string      `json:"updated_at"`
	} `json:"object_attributes"`
	ObjectKind string `json:"object_kind"`
}
