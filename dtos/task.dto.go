package dtos

type TaskDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type GetTaskDto struct {
	PageOption PageOption `json:"pageOption"`
}
