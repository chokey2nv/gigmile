package models

const (
	SettingName_StandUpStartTime = "standUpStartTime"
	SettingName_StandUpEndTime   = "standUpEndTime"
)

type Setting struct {
	Id          string `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Value       string `json:"value" bson:"value"`
	UpdateBy    string `json:"updateBy" bson:"updateBy"`
	UpdatedAt   int64  `json:"updateAt" bson:"updateAt"`
}
