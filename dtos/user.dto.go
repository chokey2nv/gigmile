package dtos

type UserDto struct {
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserRole  string `json:"userRole"`
}
type GetUserDto struct {
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
	UserRole  string `json:"userRole"`
}
type GetUsersDto struct {
	User       *GetTaskDto `json:"user"`
	PageOption PageOption  `json:"pageOption"`
}
