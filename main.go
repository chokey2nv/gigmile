package main

import "github.com/chokey2nv/gigmile/routes"

//@title Update (stand-up) API
//@version 1
//@description Create a backend API application for Gigmile Technologies' tech team to replace verbal daily standups with written updates. The application should allow employees to submit daily updates within a 15-minute window, view updates from all team members, filter updates by week/day/sprint/owner, and include specific information such as EmployeeID, Employee name, Date, Sprint ID, Task IDs, What was done yesterday, What will be done today, Blocked by employee IDs, Breakaway, Check-in time, and Status (before standup, after standup, within standup ) around what time was this logged. The system should not allow deletion or updating of updates.

//@contact.name Agu Chijioke
//@contact.url https://github.com/chokey2nv
//@contact.email chokey2nv@gmail.com

// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Initialize and run the application
	routes.Run()
}
