package controllers

import (
	"net/http"

	"github.com/chokey2nv/gigmile/config"
	"github.com/chokey2nv/gigmile/dtos"
	"github.com/chokey2nv/gigmile/services"
	"github.com/gin-gonic/gin"
)

// CreateTask handles the creation of tasks
// @Summary Create an task
// @Description Create a new task
// @Tags tasks
// @Accept json
// @Produce json
// @Param input body dtos.TaskDto true "Task information"
// @Success 200 {object} dtos.Response "Task created successfully"
// @Failure 400 {object} dtos.ErrorResponse "Bad Request"
// @Failure 500 {object} dtos.ErrorResponse "Internal Server Error"
// @Router /api/v1/tasks/new [post]
func CreateTask(c *gin.Context) {
	// Parse JSON request and save the task to MongoDB
	var taskDto dtos.TaskDto
	if err := c.BindJSON(&taskDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	taskService := services.TaskService{
		AppConfig: config.Config,
	}
	task, err := taskService.CreateTask(c, &taskDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, dtos.Response{Data: task})
}

// GetTasks handles retrieving all tasks
// @Summary Get all tasks
// @Description Get a list of all tasks
// @Tags tasks
// @Produce json
// @Param input body dtos.GetTaskDto true "User information"
// @Success 200 {array} dtos.Response "List of tasks"
// @Failure 500 {object} dtos.ErrorResponse "Internal Server Error"
// @Router /api/v1/tasks/get [post]
func GetTasks(c *gin.Context) {
	var getTaskDto dtos.GetTaskDto
	if err := c.BindJSON(&getTaskDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Retrieve all tasks from MongoDB
	taskService := services.TaskService{
		AppConfig: config.Config,
	}
	tasks, err := taskService.GetTasks(c, &getTaskDto.PageOption)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, dtos.Response{Data: tasks})
}
