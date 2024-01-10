package dtos

type SettingDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       string `json:"value"`
}
