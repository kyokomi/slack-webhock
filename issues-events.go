package main

type IssuesEvents struct {
	ObjectAttributes struct {
		AssigneeID  float64     `json:"assignee_id"`
		AuthorID    float64     `json:"author_id"`
		BranchName  interface{} `json:"branch_name"`
		CreatedAt   string      `json:"created_at"`
		Description string      `json:"description"`
		ID          float64     `json:"id"`
		Iid         float64     `json:"iid"`
		MilestoneID interface{} `json:"milestone_id"`
		Position    float64     `json:"position"`
		ProjectID   float64     `json:"project_id"`
		State       string      `json:"state"`
		Title       string      `json:"title"`
		UpdatedAt   string      `json:"updated_at"`
	} `json:"object_attributes"`
	ObjectKind string `json:"object_kind"`
}
