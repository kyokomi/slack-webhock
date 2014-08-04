package main

type IssuesEvents struct {
	ObjectAttributes struct {
		AssigneeID  int         `json:"assignee_id"`
		AuthorID    int         `json:"author_id"`
		BranchName  interface{} `json:"branch_name"`
		CreatedAt   string      `json:"created_at"`
		Description string      `json:"description"`
		ID          int         `json:"id"`
		Iid         int         `json:"iid"`
		MilestoneID int         `json:"milestone_id"`
		Position    int         `json:"position"`
		ProjectID   int         `json:"project_id"`
		State       string      `json:"state"`
		Title       string      `json:"title"`
		UpdatedAt   string      `json:"updated_at"`
	} `json:"object_attributes"`
	ObjectKind string `json:"object_kind"`
}
