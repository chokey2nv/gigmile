package controllers

import (
	"net/http"

	"github.com/chokey2nv/gigmile/config"
	"github.com/chokey2nv/gigmile/dtos"
	"github.com/chokey2nv/gigmile/services"
	"github.com/gin-gonic/gin"
)

// CreateSprint handles the creation of sprints
// @Summary Create an sprint
// @Description Create a new sprint
// @Tags sprints
// @Accept json
// @Produce json
// @Param input body models.Sprint true "Sprint information"
// @Success 200 {object} dtos.Response "Sprint created successfully"
// @Failure 400 {object} dtos.ErrorResponse "Bad Request"
// @Failure 500 {object} dtos.ErrorResponse "Internal Server Error"
// @Router /api/v1/sprints/new [post]
func CreateSprint(c *gin.Context) {
	// Parse JSON request and save the sprint to MongoDB
	var sprintDto dtos.SprintDto
	if err := c.BindJSON(&sprintDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sprintService := services.SprintService{
		AppConfig: config.Config,
	}
	sprint, err := sprintService.CreateSprint(c, &sprintDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, dtos.Response{Data: sprint})
}

// GetSprints handles retrieving all sprints
// @Summary Get all sprints
// @Description Get a list of all sprints
// @Tags sprints
// @Produce json
// @Param input body dtos.GetSprintDto true "User information"
// @Success 200 {array} dtos.Response "List of sprints"
// @Failure 500 {object} dtos.ErrorResponse "Internal Server Error"
// @Router /api/v1/sprints/get [post]
func GetSprints(c *gin.Context) {
	var getSprintsDto dtos.GetSprintsDto
	if err := c.BindJSON(&getSprintsDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Retrieve all sprints from MongoDB
	sprintService := services.SprintService{
		AppConfig: config.Config,
	}
	sprints, err := sprintService.GetSprints(c, getSprintsDto.PageOption)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, dtos.Response{Data: sprints})
}
