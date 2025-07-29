package models

type Task struct {
	ID         int    `json:"id"`
	TitleFa    string `json:"title_fa"`
	AssigneeFa string `json:"assignee_fa"`
	Start      string `json:"start"`
	End        string `json:"end"`
	Status     string `json:"status"`
	Progress   int    `json:"progress"`
}
