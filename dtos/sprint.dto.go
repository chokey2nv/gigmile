package dtos

type SprintDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}
type GetSprintDto struct {
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
type GetSprintsDto struct {
	Sprint     *GetSprintDto `json:"sprint"`
	PageOption *PageOption   `json:"pageOption"`
}
