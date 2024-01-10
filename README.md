**Gigmile Backend**

*Description*
A backend API application for Gigmile Technologies' tech team to replace verbal daily standups with written updates. The application allows employees to submit daily updates within a 15-minute window, view updates from all team members, filter updates by 
- week
- dayOfWeek
- sprint name
- employee name
- employee id
  
and include specific information such as EmployeeID, Employee name, Date, Sprint ID, Task IDs, What was done yesterday, What will be done today, Blocked by employee IDs, Breakaway, Check-in time, and Status (before standup, after standup, within standup ) around what time was this logged. The system should not allow deletion or updating of updates by simply not providing the endpoint.

Here is the return data structure 
```go
type UpdateUserResponseDto struct {
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
	CreateAt                 primitive.DateTime     `json:"createdAt"`
}
```
**Flow**

- Signup Admin 
- Signup staff 
- Admin creates settings (stand-up start and end time)
- Admin creates sprint
- Admin create tasks 
- staff or admin can use the sprint and tasks created by admin to register a new update 

`Note - Authentication and access_token is implemented for this endpoint, hence signup and login is required to access the rest of the endpoints. On signup or login, an access token is provided which should be registered as a request header - Authorization`.

**Folder Structure**

```
sh
-app-
  |
  |- config: contains app configurations/setup
  | - controllers: the endpoint interfaces
  | - dtos: data transfer objects for the interface
  | - models: data models for the database
  | - routes: this contains a single router file that defines all the paths
  | - services: this is the internal logics and communication to database
  | - main.go: entry file
```
**Database**

This application made use of mongodb database. Ensure that your mongodb instance is up before running this application. 

If you don't have database in your system, no worries, there is an option for you. If you don't like loading additional softwares on your system and only goes the `docker` way, we gat you covered. Find the docker setup below.

**Setup**

Clone this repo and run 
```sh
go mod download
```
or
```sh
go mod tidy
```

**Run Test**

This application uses test suits and achieved via ginkgo. Ginkgo bootstrap is already ran. Find more information here https://onsi.github.io/ginkgo/

install ginkgo in your system using the command 
```sh 
go install github.com/onsi/ginkgo/v2/ginkgo
```
the ginkgo lib for the project is already added when you run the setup command above, but if still having an issue with that, you can add it directly by following the link above.

Run the test with the command line
```sh
ginkgo -v
```

**Run Application**

Run this command line from the root folder
```sh
go run main.go
``` 
The application runs on port 8080 and versioned. The current version is version 1 (v1). The entry path is /api.

`http://localhost:8080/api/v1`

*Routes* 
- signup 
- login
- users
- settings
- sprints
- tasks
- updates 

`eg - http://localhost:8080/api/v1/signup`



**Documentation**

This application has `swagger` integration, and doc already generated `/docs`. 
Run the application and find the swagger documentation in the `/swagger/index.html` of your host in your browser.
http://localhost:8080/swagger/index.html 

this is a public path and isn't locked by authentication. 

**Docker**

If you have docker installed in your system, you can simply run this app using 
```sh
docker-compose up
```