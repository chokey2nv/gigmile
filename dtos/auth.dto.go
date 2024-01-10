package dtos

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpDto struct {
	User     UserDto `json:"user"`
}
