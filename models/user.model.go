package models

type UserRole = string

const (
	UserRole_Admin UserRole = "admin"
	UserRole_Staff UserRole = "staff"
)

// User represents the structure of a user
type User struct {
	Id        string   `json:"id" bson:"id"`
	LastName  string   `json:"lastName" bson:"lastName"`
	FirstName string   `json:"firstName" bson:"firstName"`
	UserRole  UserRole `json:"userRole" bson:"userRole"`
	Email     string   `json:"email" bson:"email"`
	Hash      string   `json:"hash" bson:"hash"`
}
