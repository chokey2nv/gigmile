package dtos

import (
	"github.com/chokey2nv/gigmile/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUserResponseDto struct {
	Id        string          `json:"id"`
	LastName  string          `json:"lastName"`
	FirstName string          `json:"firstName"`
	UserRole  models.UserRole `json:"userRole"`
	Email     string          `json:"email"`
}
type UpdateResponseDto struct {
	Id                       string                 `json:"id"`
	Employee                 *UpdateUserResponseDto `json:"employee"`
	EmployeeID               string                 `json:"employeeId"`
	Timestamp                int64                  `json:"timestamp"`
	Week                     int64                  `json:"week"`
	DayOfWeek                int64                  `json:"dayOfWeek"`
	SprintID                 string                 `json:"sprintId"`
	TaskIDs                  []string               `json:"taskIds"`
	PreviousCompletedTaskIDs []string               `json:"previousCompletedTaskIds"`
	CurrentTaskIDs           []string               `json:"currentTaskIds"`
	BlockedByEmployeeIDs     []string               `json:"blockByEmployeeIds"`
	Breakaway                bool                   `json:"breakaway"`
	CheckedInTime            string                 `json:"checkedInTime"`
	Status                   string                 `json:"status"`
	CreateAt                 primitive.DateTime     `json:"createdAt" swaggertype:"primitive.DateTime" format:"string"`
}
type UpdateDto struct {
	SprintID                 string   `json:"sprintId"`
	Week                     int64    `json:"week"`
	TaskIDs                  []string `json:"taskIds"`
	PreviousCompletedTaskIDs []string `json:"previousCompletedTaskIds"`
	CurrentTaskIDs           []string `json:"currentTaskIds"`
	BlockedByEmployeeIDs     []string `json:"blockByEmployeeIds"`
	Breakaway                bool     `json:"breakaway"`
}
type GetUPdateFilterDto struct {
	//week/day/sprint/owner
	EmployeeID   string `json:"employeeId"`
	EmployeeName string `json:"employeeName"`
	Week         int64  `json:"week"`
	DayOfWeek    int64  `json:"dayOfWeek"`
	Date         string `json:"date"`
	SprintName   string `json:"sprintName"`
}
type GetUpdateDto struct {
	Filter     *GetUPdateFilterDto `json:"filter"`
	PageOption *PageOption         `json:"pageOption"`
}
