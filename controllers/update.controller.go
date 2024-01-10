package controllers

import (
	"log"
	"net/http"

	"github.com/chokey2nv/gigmile/config"
	"github.com/chokey2nv/gigmile/dtos"
	"github.com/chokey2nv/gigmile/services"
	"github.com/gin-gonic/gin"
)

// CreateUpdate handles the creation of updates
// @Summary Create an update
// @Description Create a new update
// @Tags updates
// @Accept json
// @Produce json
// @Param input body dtos.UpdateDto true "Update information"
// @Success 200 {object} dtos.Response "Update created successfully"
// @Failure 400 {object} dtos.ErrorResponse "Bad Request"
// @Failure 500 {object} dtos.ErrorResponse "Internal Server Error"
// @Router /api/v1/updates/new [post]
func CreateUpdate(c *gin.Context) {
	// Parse JSON request and save the update to MongoDB
	var updateDto dtos.UpdateDto
	if err := c.BindJSON(&updateDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateService := &services.UpdateService{
		AppConfig: config.Config,
	}
	update, err := updateService.CreateUpdate(c, &updateDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dtos.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, dtos.Response{Data: &update})
}

// GetUpdates handles retrieving all updates
// @Summary Get all updates
// @Description Get a list of all updates
// @Tags updates
// @Produce json
// @Param input body dtos.GetUpdateDto true "Update information"
// @Success 200 {array} []models.Update "List of updates"
// @Failure 500 {object} dtos.ErrorResponse "Internal Server Error"
// @Router /api/v1/updates/get [post]
func GetUpdates(c *gin.Context) {
	getUpdateDto := dtos.GetUpdateDto{}
	if err := c.BindJSON(&getUpdateDto); err != nil {
		log.Println(err)
	}
	// Retrieve all updates from MongoDB
	updateService := &services.UpdateService{
		AppConfig: config.Config,
	}
	updates, err := updateService.GetUpdates(c, getUpdateDto.Filter, getUpdateDto.PageOption)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtos.Response{Data: &updates})
}
