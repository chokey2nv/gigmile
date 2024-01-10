// models/update_model.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Update represents the update data structure
type Update struct {
	Id                       string             `json:"id" bson:"id"`
	EmployeeID               string             `json:"employeeId" bson:"employeeId"`
	Timestamp                int64              `json:"timestamp" bson:"timestamp"`
	Week                     int64              `json:"week" bson:"week"`
	DayOfWeek                int64              `json:"dayOfWeek" bson:"dayOfWeek"`
	SprintID                 string             `json:"sprintId" bson:"sprintId"`
	TaskIDs                  []string           `json:"taskIds" bson:"taskIds"`
	PreviousCompletedTaskIDs []string           `json:"previousCompletedTaskIds" bson:"previousCompletedTaskIds"`
	CurrentTaskIDs           []string           `json:"currentTaskIds" bson:"currentTaskIds"`
	BlockedByEmployeeIDs     []string           `json:"blockByEmployeeIds" bson:"blockByEmployeeIds"`
	Breakaway                bool               `json:"breakaway" bson:"Breakaway"`
	Employee                 *User              `json:"employee" bson:"employee"`
	CreateAt                 primitive.DateTime `json:"createdAt" bson:"createdAt" swaggertype:"string"`
}
