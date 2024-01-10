package models

type Sprint struct {
	Id             string `json:"id" bson:"id"`
	Name           string `json:"name" bson:"name"`
	Description    string `json:"description" bson:"description"`
	StartTimestamp int64  `json:"startTimestamp" bson:"startTimestamp"`
	EndTimestamp   int64  `json:"endTimestamp" bson:"endTimestamp"`
}
