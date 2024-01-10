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

**Highlights**
Note that settings can not be any name, it must be one of the following strings as defined in the constants

```go
const (
	SettingName_StandUpStartTime = "standUpStartTime"
	SettingName_StandUpEndTime   = "standUpEndTime"
)
```
`standUpStartTime` and `standUpEndTime`


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

**Sample Flow / Api Calls**

*Sign up*

Request 
```sh
curl -X 'POST' \
  'http://localhost:8080/api/v1/signup' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "user": {
    "email": "chokey2nv@gmail.com",
    "firstName": "Chijioke",
    "lastName": "Agu",
    "password": "password",
    "userRole": "admin"
  }
}'
```
Response body 
```sh
{
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NzU2ODcsInVzZXIiOnsiaWQiOiJhMGU5MGJlOC1jOTcxLTQzZGYtYjY2Yy1iMWVkMWNiZTJhYTUiLCJsYXN0TmFtZSI6IkFndSIsImZpcnN0TmFtZSI6IkNoaWppb2tlIiwidXNlclJvbGUiOiJhZG1pbiIsImVtYWlsIjoiY2hva2V5Mm52QGdtYWlsLmNvbSIsImhhc2giOiIkMmEkMTAkZUluUFpQNm02VEpaT2gubXZ1NDdpZXcwS3lGRXRxSnV4NGx5djdrVUhLS25KOGFsOEltc0MifX0.0hqAyrVTlNj6FnaQs4TTyiqXA6zotLLQsIbERW2Pr6A",
    "user": {
      "id": "a0e90be8-c971-43df-b66c-b1ed1cbe2aa5",
      "lastName": "Agu",
      "firstName": "Chijioke",
      "userRole": "admin",
      "email": "chokey2nv@gmail.com"
    }
  }
}
```

*Sign in*

Request 
```sh
curl -X 'POST' \
  'http://localhost:8080/api/v1/login' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "chokey2nv@gmail.com",
  "password": "password"
}'
```
Response 
```sh
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NzYxNDcsInVzZXIiOnsiaWQiOiJhMGU5MGJlOC1jOTcxLTQzZGYtYjY2Yy1iMWVkMWNiZTJhYTUiLCJsYXN0TmFtZSI6IkFndSIsImZpcnN0TmFtZSI6IkNoaWppb2tlIiwidXNlclJvbGUiOiJhZG1pbiIsImVtYWlsIjoiY2hva2V5Mm52QGdtYWlsLmNvbSIsImhhc2giOiIkMmEkMTAkZUluUFpQNm02VEpaT2gubXZ1NDdpZXcwS3lGRXRxSnV4NGx5djdrVUhLS25KOGFsOEltc0MifX0.ZiVFmI1zJaEjSlctantBrUvw8JYWiZhk0imHH9KfXP4"
}
```
Response this access_token with yours.

*Settings*

Create Stand-up start time 
```sh
curl --location --request POST 'localhost:8080/api/v1/settings' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NzY3MDAsInVzZXIiOnsiaWQiOiJhMGU5MGJlOC1jOTcxLTQzZGYtYjY2Yy1iMWVkMWNiZTJhYTUiLCJsYXN0TmFtZSI6IkFndSIsImZpcnN0TmFtZSI6IkNoaWppb2tlIiwidXNlclJvbGUiOiJhZG1pbiIsImVtYWlsIjoiY2hva2V5Mm52QGdtYWlsLmNvbSIsImhhc2giOiIkMmEkMTAkZUluUFpQNm02VEpaT2gubXZ1NDdpZXcwS3lGRXRxSnV4NGx5djdrVUhLS25KOGFsOEltc0MifX0.g5WgdjZDej4OZT-mfsnbPRna_e1zanNUZZG8GLZ1UFY' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name" : "standUpStartTime",
    "value": "15:00",
    "description": "Stand-up start time"
}'
```

Response 
```sh
{
    "data": {
        "id": "30e5c889-43ff-4354-9d91-735e0322ed92",
        "name": "standUpStartTime",
        "description": "Stand-up start time",
        "value": "15:00",
        "updateBy": "a0e90be8-c971-43df-b66c-b1ed1cbe2aa5",
        "updateAt": 1704873549410
    }
}
```

*Create Stand up end time*

Request
```sh
curl --location --request POST 'localhost:8080/api/v1/settings' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NzY3MDAsInVzZXIiOnsiaWQiOiJhMGU5MGJlOC1jOTcxLTQzZGYtYjY2Yy1iMWVkMWNiZTJhYTUiLCJsYXN0TmFtZSI6IkFndSIsImZpcnN0TmFtZSI6IkNoaWppb2tlIiwidXNlclJvbGUiOiJhZG1pbiIsImVtYWlsIjoiY2hva2V5Mm52QGdtYWlsLmNvbSIsImhhc2giOiIkMmEkMTAkZUluUFpQNm02VEpaT2gubXZ1NDdpZXcwS3lGRXRxSnV4NGx5djdrVUhLS25KOGFsOEltc0MifX0.g5WgdjZDej4OZT-mfsnbPRna_e1zanNUZZG8GLZ1UFY' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name" : "standUpEndTime",
    "value": "15:15",
    "description": "Stand-up end time"
}'
```

Response
```
{
    "data": {
        "id": "cd0853ec-ee15-4952-b86b-eabcb6aadc0a",
        "name": "standUpEndTime",
        "description": "Stand-up end time",
        "value": "15:15",
        "updateBy": "a0e90be8-c971-43df-b66c-b1ed1cbe2aa5",
        "updateAt": 1704873710660
    }
}
```

*Create Sprint*

Request 
```sh
curl --location --request POST 'localhost:8080/api/v1/sprints/new' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NzY3MDAsInVzZXIiOnsiaWQiOiJhMGU5MGJlOC1jOTcxLTQzZGYtYjY2Yy1iMWVkMWNiZTJhYTUiLCJsYXN0TmFtZSI6IkFndSIsImZpcnN0TmFtZSI6IkNoaWppb2tlIiwidXNlclJvbGUiOiJhZG1pbiIsImVtYWlsIjoiY2hva2V5Mm52QGdtYWlsLmNvbSIsImhhc2giOiIkMmEkMTAkZUluUFpQNm02VEpaT2gubXZ1NDdpZXcwS3lGRXRxSnV4NGx5djdrVUhLS25KOGFsOEltc0MifX0.g5WgdjZDej4OZT-mfsnbPRna_e1zanNUZZG8GLZ1UFY' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name" : "Q1 Sprint - Week 1",
    "description": "This is the first sprint (week 1) of the first quarter of the year",
    "startDate":   "2024-01-01",
    "endDate":     "2022-01-08"
}'
```
Response 
```sh
{
    "data": {
        "id": "c2f71d8c-0083-45de-87fb-17e3e55c9e39",
        "name": "Q1 Sprint - Week 1",
        "description": "This is the first sprint (week 1) of the first quarter of the year",
        "startTimestamp": 1704067200000,
        "endTimestamp": 1659312000000
    }
}
```

*Create Task*
Request 

```sh
curl --location --request POST 'localhost:8080/api/v1/tasks/new' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NzY3MDAsInVzZXIiOnsiaWQiOiJhMGU5MGJlOC1jOTcxLTQzZGYtYjY2Yy1iMWVkMWNiZTJhYTUiLCJsYXN0TmFtZSI6IkFndSIsImZpcnN0TmFtZSI6IkNoaWppb2tlIiwidXNlclJvbGUiOiJhZG1pbiIsImVtYWlsIjoiY2hva2V5Mm52QGdtYWlsLmNvbSIsImhhc2giOiIkMmEkMTAkZUluUFpQNm02VEpaT2gubXZ1NDdpZXcwS3lGRXRxSnV4NGx5djdrVUhLS25KOGFsOEltc0MifX0.g5WgdjZDej4OZT-mfsnbPRna_e1zanNUZZG8GLZ1UFY' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name" : "Build Login Endpoint",
    "description": "This is task is to build a login endpoint for the front-end to consume"
}'
```
Response 
```sh
{
    "data": {
        "id": "ee31b80d-d9a1-4d83-81ae-c7867ef476d3",
        "name": "Build Login Endpoint",
        "description": "This is task is to build a login endpoint for the front-end to consume",
        "createdBy": "a0e90be8-c971-43df-b66c-b1ed1cbe2aa5",
        "createAt": 1704874208827
    }
}
```

*Create Update (stand-up)*
Request 
```sh
curl --location --request POST 'localhost:8080/api/v1/updates/new' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NzY3MDAsInVzZXIiOnsiaWQiOiJhMGU5MGJlOC1jOTcxLTQzZGYtYjY2Yy1iMWVkMWNiZTJhYTUiLCJsYXN0TmFtZSI6IkFndSIsImZpcnN0TmFtZSI6IkNoaWppb2tlIiwidXNlclJvbGUiOiJhZG1pbiIsImVtYWlsIjoiY2hva2V5Mm52QGdtYWlsLmNvbSIsImhhc2giOiIkMmEkMTAkZUluUFpQNm02VEpaT2gubXZ1NDdpZXcwS3lGRXRxSnV4NGx5djdrVUhLS25KOGFsOEltc0MifX0.g5WgdjZDej4OZT-mfsnbPRna_e1zanNUZZG8GLZ1UFY' \
--header 'Content-Type: application/json' \
--data-raw '{
    "week": 1,
    "sprintId" : "c2f71d8c-0083-45de-87fb-17e3e55c9e39",
    "taskIds": ["ee31b80d-d9a1-4d83-81ae-c7867ef476d3"],
    "previousCompletedTaskIds": ["ee31b80d-d9a1-4d83-81ae-c7867ef476d3"],
    "currentTaskIds": ["ee31b80d-d9a1-4d83-81ae-c7867ef476d3"],
    "blockByEmployeeIds": [],
    "breakaway": false
}'
```
Response 

```sh
{
    "data": {
        "id": "4d1268c1-3045-4b85-840d-b22d75d28f2c",
        "employeeId": "a0e90be8-c971-43df-b66c-b1ed1cbe2aa5",
        "timestamp": 1704874867804,
        "week": 1,
        "dayOfWeek": 0,
        "sprintId": "c2f71d8c-0083-45de-87fb-17e3e55c9e39",
        "taskIds": [
            "ee31b80d-d9a1-4d83-81ae-c7867ef476d3"
        ],
        "previousCompletedTaskIds": [
            "ee31b80d-d9a1-4d83-81ae-c7867ef476d3"
        ],
        "currentTaskIds": [
            "ee31b80d-d9a1-4d83-81ae-c7867ef476d3"
        ],
        "blockByEmployeeIds": [],
        "breakaway": false,
        "employee": null,
        "createdAt": "2024-01-10T08:21:07.804Z"
    }
}
```

*Get All Updates*

Request 
```sh
curl --location --request POST 'localhost:8080/api/v1/updates/get' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NzY3MDAsInVzZXIiOnsiaWQiOiJhMGU5MGJlOC1jOTcxLTQzZGYtYjY2Yy1iMWVkMWNiZTJhYTUiLCJsYXN0TmFtZSI6IkFndSIsImZpcnN0TmFtZSI6IkNoaWppb2tlIiwidXNlclJvbGUiOiJhZG1pbiIsImVtYWlsIjoiY2hva2V5Mm52QGdtYWlsLmNvbSIsImhhc2giOiIkMmEkMTAkZUluUFpQNm02VEpaT2gubXZ1NDdpZXcwS3lGRXRxSnV4NGx5djdrVUhLS25KOGFsOEltc0MifX0.g5WgdjZDej4OZT-mfsnbPRna_e1zanNUZZG8GLZ1UFY' \
--header 'Content-Type: application/json' \
--data-raw '{}'
```

Response 
```sh
{
    "data": [
        {
            "id": "4d1268c1-3045-4b85-840d-b22d75d28f2c",
            "employee": {
                "id": "a0e90be8-c971-43df-b66c-b1ed1cbe2aa5",
                "lastName": "Agu",
                "firstName": "Chijioke",
                "userRole": "admin",
                "email": "chokey2nv@gmail.com"
            },
            "employeeId": "a0e90be8-c971-43df-b66c-b1ed1cbe2aa5",
            "timestamp": 1704874867804,
            "week": 1,
            "dayOfWeek": 4,
            "sprintId": "c2f71d8c-0083-45de-87fb-17e3e55c9e39",
            "taskIds": [
                "ee31b80d-d9a1-4d83-81ae-c7867ef476d3"
            ],
            "previousCompletedTaskIds": [
                "ee31b80d-d9a1-4d83-81ae-c7867ef476d3"
            ],
            "currentTaskIds": [
                "ee31b80d-d9a1-4d83-81ae-c7867ef476d3"
            ],
            "blockByEmployeeIds": [],
            "breakaway": false,
            "checkedInTime": "09:21",
            "status": "Before stand-up",
            "createdAt": "2024-01-10T08:21:07.804Z"
        }
    ]
}
```

`You can fetch using several filters as below`

*By Employee Name*
```sh
curl --location --request POST 'localhost:8080/api/v1/updates/get' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NzY3MDAsInVzZXIiOnsiaWQiOiJhMGU5MGJlOC1jOTcxLTQzZGYtYjY2Yy1iMWVkMWNiZTJhYTUiLCJsYXN0TmFtZSI6IkFndSIsImZpcnN0TmFtZSI6IkNoaWppb2tlIiwidXNlclJvbGUiOiJhZG1pbiIsImVtYWlsIjoiY2hva2V5Mm52QGdtYWlsLmNvbSIsImhhc2giOiIkMmEkMTAkZUluUFpQNm02VEpaT2gubXZ1NDdpZXcwS3lGRXRxSnV4NGx5djdrVUhLS25KOGFsOEltc0MifX0.g5WgdjZDej4OZT-mfsnbPRna_e1zanNUZZG8GLZ1UFY' \
--header 'Content-Type: application/json' \
--data-raw '{
    "filter": {
        "employeeName": "Agu"
    }
}'
```

*By Employee ID*
```sh
curl --location --request POST 'localhost:8080/api/v1/updates/get' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NzY3MDAsInVzZXIiOnsiaWQiOiJhMGU5MGJlOC1jOTcxLTQzZGYtYjY2Yy1iMWVkMWNiZTJhYTUiLCJsYXN0TmFtZSI6IkFndSIsImZpcnN0TmFtZSI6IkNoaWppb2tlIiwidXNlclJvbGUiOiJhZG1pbiIsImVtYWlsIjoiY2hva2V5Mm52QGdtYWlsLmNvbSIsImhhc2giOiIkMmEkMTAkZUluUFpQNm02VEpaT2gubXZ1NDdpZXcwS3lGRXRxSnV4NGx5djdrVUhLS25KOGFsOEltc0MifX0.g5WgdjZDej4OZT-mfsnbPRna_e1zanNUZZG8GLZ1UFY' \
--header 'Content-Type: application/json' \
--data-raw '{
    "filter": {
        "employeeId": "a0e90be8-c971-43df-b66c-b1ed1cbe2aa5"
    }
}'
```
*By Week*
```sh
curl --location --request POST 'localhost:8080/api/v1/updates/get' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NzY3MDAsInVzZXIiOnsiaWQiOiJhMGU5MGJlOC1jOTcxLTQzZGYtYjY2Yy1iMWVkMWNiZTJhYTUiLCJsYXN0TmFtZSI6IkFndSIsImZpcnN0TmFtZSI6IkNoaWppb2tlIiwidXNlclJvbGUiOiJhZG1pbiIsImVtYWlsIjoiY2hva2V5Mm52QGdtYWlsLmNvbSIsImhhc2giOiIkMmEkMTAkZUluUFpQNm02VEpaT2gubXZ1NDdpZXcwS3lGRXRxSnV4NGx5djdrVUhLS25KOGFsOEltc0MifX0.g5WgdjZDej4OZT-mfsnbPRna_e1zanNUZZG8GLZ1UFY' \
--header 'Content-Type: application/json' \
--data-raw '{
    "filter": {
        "week": 1
    }
}'
```
*By DayOfWeek*
From our results so far you can see that day of week is `4` - Thursday. So can search for such or any other 
```sh
curl --location --request POST 'localhost:8080/api/v1/updates/get' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NzY3MDAsInVzZXIiOnsiaWQiOiJhMGU5MGJlOC1jOTcxLTQzZGYtYjY2Yy1iMWVkMWNiZTJhYTUiLCJsYXN0TmFtZSI6IkFndSIsImZpcnN0TmFtZSI6IkNoaWppb2tlIiwidXNlclJvbGUiOiJhZG1pbiIsImVtYWlsIjoiY2hva2V5Mm52QGdtYWlsLmNvbSIsImhhc2giOiIkMmEkMTAkZUluUFpQNm02VEpaT2gubXZ1NDdpZXcwS3lGRXRxSnV4NGx5djdrVUhLS25KOGFsOEltc0MifX0.g5WgdjZDej4OZT-mfsnbPRna_e1zanNUZZG8GLZ1UFY' \
--header 'Content-Type: application/json' \
--data-raw '{
    "filter": {
        "dayOfWeek": 1
    }
}'
```