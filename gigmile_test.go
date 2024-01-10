package main_test

import (
	"context"

	"github.com/chokey2nv/gigmile/config"
	"github.com/chokey2nv/gigmile/dtos"
	"github.com/chokey2nv/gigmile/models"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var errorf = config.Errorf

type SignUpData struct {
	User        *models.User `json:"user"`
	AccessToken string       `json:"access_token"`
}
type Response struct {
	Data    any    `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type Account struct {
	id          string
	email       string
	password    string
	accessToken string
	firstName   string
	lastName    string
}

var _ = Describe("Ordered", Ordered, func() {
	var (
		req    *resty.Client
		resp   *resty.Response
		err    error
		result *Response
		admin  = Account{
			email:       "admin@gmail.com",
			password:    "adminPass",
			accessToken: "",
			id:          "",
			lastName:    "Admin",
			firstName:   "Chijioke",
		}
		staff = Account{
			email:       "staff@gmail.com",
			password:    "staffPass",
			accessToken: "",
			id:          "",
			lastName:    "Staff",
			firstName:   "Chima",
		}
		sprintID    string
		taskID      string
		currentWeek = 1
		sprintName  = "Q1 Sprint 1"
	)
	BeforeAll(func() {
		req = resty.New().SetBaseURL("http://localhost:8080/api/v1")
	})
	Describe("Sign up", func() {
		Context("POST /signup", func() {
			Describe("Create Admin Account", func() {
				It("should return status 200 and expected response body", func() {
					resp, err = req.R().SetBody(dtos.SignUpDto{
						User: dtos.UserDto{
							LastName:  admin.lastName,
							FirstName: admin.firstName,
							Email:     admin.email,
							Password:  admin.password,
							UserRole:  "admin",
						},
					}).SetResult(&result).Post("/signup")
					if err != nil {
						panic(err)
					}
					Expect(result.Data).NotTo(Equal(""))
					Expect(resp.StatusCode()).To(Equal(200))
					signUpData := SignUpData{}
					config.ToJSONStruct(result.Data, &signUpData)
					admin.accessToken = signUpData.AccessToken
					admin.id = signUpData.User.Id
				})
			})
			Describe("Create Staff Account", func() {
				It("should return status 200 and expected response body", func() {
					resp, err = req.R().SetBody(dtos.SignUpDto{
						User: dtos.UserDto{
							LastName:  staff.lastName,
							FirstName: staff.firstName,
							Email:     staff.email,
							Password:  staff.password,
							UserRole:  "admin",
						},
					}).SetResult(&result).Post("/signup")
					if err != nil {
						panic(err)
					}
					Expect(result.Data).NotTo(Equal(""))
					Expect(resp.StatusCode()).To(Equal(200))
					signUpData := SignUpData{}
					config.ToJSONStruct(result.Data, &signUpData)
					staff.accessToken = signUpData.AccessToken
					staff.id = signUpData.User.Id
				})
			})
		})
	})
	Describe("Login Admin", func() {
		Context("POST /login", func() {
			It("should return status 200 and expected response body", func() {
				resp, err = req.R().SetBody(dtos.LoginDto{
					Email:    admin.email,
					Password: admin.password,
				}).SetResult(&result).Post("/login")
				if err != nil {
					panic(err)
				}
				Expect(result.Data).NotTo(Equal(""))
				Expect(resp.StatusCode()).To(Equal(200))
				loginData := SignUpData{}
				config.ToJSONStruct(result.Data, &loginData)
				admin.accessToken = loginData.AccessToken
			})
		})
	})
	Describe("Admin Create Setting - Stand-up start time", func() {
		Context("POST /settings", func() {
			It("should return status 201 and expected response body", func() {
				// var result SignUpResult
				resp, err = req.R().SetBody(dtos.SettingDto{
					Name:        models.SettingName_StandUpStartTime,
					Description: "This is stand-up start time",
					Value:       "15:00",
				}).SetResult(&result).SetHeader("Authorization", admin.accessToken).Post("/settings")
				if err != nil {
					panic(err)
				}
				Expect(result.Data).NotTo(Equal(""))
				Expect(resp.StatusCode()).To(Equal(201))
			})
		})
	})
	Describe("Admin Create Setting - Stand-up end time", func() {
		Context("POST /settings", func() {
			It("should return status 201 and expected response body", func() {
				// var result SignUpResult
				resp, err = req.R().SetBody(dtos.SettingDto{
					Name:        models.SettingName_StandUpEndTime,
					Description: "This is stand-up end time",
					Value:       "15:30",
				}).SetResult(&result).SetHeader("Authorization", admin.accessToken).Post("/settings")
				if err != nil {
					panic(err)
				}
				Expect(result.Data).NotTo(Equal(""))
				Expect(resp.StatusCode()).To(Equal(201))
			})
		})
	})
	Describe("Admin Create Sprint", func() {
		Context("POST /sprints/new", func() {
			It("should return status 200 and expected response body", func() {
				// var result SignUpResult
				resp, err = req.R().SetBody(dtos.SprintDto{
					Name:        sprintName,
					Description: "This is the first sprint (week 1) of the first quarter of the year",
					StartDate:   "2024-01-01",
					EndDate:     "2022-01-08",
				}).SetResult(&result).SetHeader("Authorization", admin.accessToken).Post("/sprints/new")

				if err != nil {
					panic(err)
				}
				Expect(result.Data).NotTo(Equal(""))
				Expect(resp.StatusCode()).To(Equal(201))
				sprint := models.Sprint{}
				config.ToJSONStruct(result.Data, &sprint)
				sprintID = sprint.Id
			})
		})
	})
	Describe("Admin Create Task", func() {
		Context("POST /tasks/new", func() {
			It("should return status 201 and expected response body", func() {
				// var result SignUpResult
				resp, err = req.R().SetBody(dtos.TaskDto{
					Name:        "Build Login Endpoint",
					Description: "This is task is to build a login endpoint for the front-end to consume",
				}).SetResult(&result).SetHeader("Authorization", admin.accessToken).Post("/tasks/new")

				if err != nil {
					panic(err)
				}
				Expect(result.Data).NotTo(Equal(""))
				Expect(resp.StatusCode()).To(Equal(201))
				task := models.Task{}
				config.ToJSONStruct(result.Data, &task)
				taskID = task.Id
			})
		})
	})
	Describe("Staff1 Create Update (stand-up)", func() {
		Context("POST /updates/new", func() {
			It("should return status 201 and expected response body", func() {
				// var result SignUpResult
				resp, err = req.R().SetBody(dtos.UpdateDto{
					SprintID:                 sprintID,
					TaskIDs:                  []string{taskID},
					PreviousCompletedTaskIDs: []string{taskID},
					CurrentTaskIDs:           []string{taskID},
					BlockedByEmployeeIDs:     []string{admin.id},
					Breakaway:                false,
					Week:                     int64(currentWeek),
				}).SetResult(&result).SetHeader("Authorization", staff.accessToken).Post("/updates/new")
				if err != nil {
					panic(err)
				}
				Expect(result.Data).NotTo(Equal(""))
				Expect(resp.StatusCode()).To(Equal(201))
			})
		})
	})
	Describe("Get Updates", func() {
		Context("POST /updates/get", func() {
			Describe("Get All Updates", func() {
				It("should return status 200 and expected response body", func() {
					resp, err = req.R().
						SetBody(dtos.GetUpdateDto{}).
						SetResult(&result).
						SetHeader("Authorization", staff.accessToken).
						Post("/updates/get")

					if err != nil {
						panic(err)
					}
					Expect(result.Data).NotTo(Equal(""))
					Expect(resp.StatusCode()).To(Equal(200))
					updateResult := []*dtos.UpdateResponseDto{}
					config.ToJSONStruct(result.Data, &updateResult)
					Expect(len(updateResult)).To(Equal(1))
				})
			})
			Describe("Get Updates By Owner - EmployeeID", func() {
				It("should return status 200 and expected response body", func() {
					resp, err = req.R().
						SetBody(dtos.GetUpdateDto{
							Filter: &dtos.GetUPdateFilterDto{
								EmployeeID: staff.id,
							},
						}).
						SetResult(&result).
						SetHeader("Authorization", staff.accessToken).
						Post("/updates/get")
					if err != nil {
						panic(err)
					}
					Expect(result.Data).NotTo(Equal(""))
					Expect(resp.StatusCode()).To(Equal(200))
					updateResult := []*dtos.UpdateResponseDto{}
					config.ToJSONStruct(result.Data, &updateResult)
					Expect(len(updateResult)).To(Equal(1))
					errorf(updateResult[0])
				})
			})
			Describe("Get Updates By Owner - Employee name", func() {
				It("should return status 200 and expected response body", func() {
					resp, err = req.R().
						SetBody(dtos.GetUpdateDto{
							Filter: &dtos.GetUPdateFilterDto{
								EmployeeName: staff.lastName,
							},
						}).
						SetResult(&result).
						SetHeader("Authorization", staff.accessToken).
						Post("/updates/get")
					if err != nil {
						panic(err)
					}
					Expect(result.Data).NotTo(Equal(""))
					Expect(resp.StatusCode()).To(Equal(200))
					updateResult := []*dtos.UpdateResponseDto{}
					config.ToJSONStruct(result.Data, &updateResult)
					Expect(len(updateResult)).To(Equal(1))
				})
			})
			Describe("Get Updates By Week", func() {
				It("should return status 200 and expected response body", func() {
					resp, err = req.R().
						SetBody(dtos.GetUpdateDto{
							Filter: &dtos.GetUPdateFilterDto{
								Week: int64(currentWeek),
							},
						}).
						SetResult(&result).
						SetHeader("Authorization", staff.accessToken).
						Post("/updates/get")

					if err != nil {
						panic(err)
					}
					Expect(result.Data).NotTo(Equal(""))
					Expect(resp.StatusCode()).To(Equal(200))
					updateResult := []*dtos.UpdateResponseDto{}
					config.ToJSONStruct(result.Data, &updateResult)
					Expect(len(updateResult)).To(Equal(1))
					config.Errorf(updateResult[0].Week)
					config.Errorf(updateResult[0].DayOfWeek)
				})
			})
			Describe("Get Updates By Sprint name", func() {
				It("should return status 200 and expected response body", func() {
					resp, err = req.R().
						SetBody(dtos.GetUpdateDto{
							Filter: &dtos.GetUPdateFilterDto{
								SprintName: sprintName,
							},
						}).
						SetResult(&result).
						SetHeader("Authorization", staff.accessToken).
						Post("/updates/get")

					if err != nil {
						panic(err)
					}
					Expect(result.Data).NotTo(Equal(""))
					Expect(resp.StatusCode()).To(Equal(200))
					updateResult := []*dtos.UpdateResponseDto{}
					config.ToJSONStruct(result.Data, &updateResult)
					Expect(len(updateResult)).To(Equal(1))
				})
			})

		})
	})
	AfterAll(func() {
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			panic(err)
		}
		client.Database("gigmile").Drop(context.Background())
	})
})
