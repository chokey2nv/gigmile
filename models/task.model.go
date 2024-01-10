package models

type Task struct {
	Id          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	CreateBy    string `json:"createdBy" bson:"createdBy"`
	CreateAt    int64  `json:"createAt" bson:"createAt"`
}
