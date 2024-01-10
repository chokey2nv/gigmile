package controllers

import (
	"net/http"

	"github.com/chokey2nv/gigmile/config"
	"github.com/chokey2nv/gigmile/dtos"
	"github.com/chokey2nv/gigmile/services"
	"github.com/gin-gonic/gin"
)

// CreateSetting handles the creation of settings
// @Summary Create an setting
// @Description Create a new setting
// @Tags settings
// @Accept json
// @Produce json
// @Param input body models.Setting true "Setting information"
// @Success 200 {object} dtos.Response "Setting created successfully"
// @Failure 400 {object} dtos.ErrorResponse "Bad Request"
// @Failure 500 {object} dtos.ErrorResponse "Internal Server Error"
// @Router /api/v1/settings [post]
func CreateSetting(c *gin.Context) {
	// Parse JSON request and save the setting to MongoDB
	var settingDto dtos.SettingDto
	if err := c.BindJSON(&settingDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	settingService := services.SettingService{
		AppConfig: config.Config,
	}
	setting, err := settingService.CreateSetting(c, &settingDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, dtos.Response{Data: setting})
}

// GetSettings handles retrieving all settings
// @Summary Get all settings
// @Description Get a list of all settings
// @Tags settings
// @Produce json
// @Success 200 {array} dtos.Response "List of settings"
// @Failure 500 {object} dtos.ErrorResponse "Internal Server Error"
// @Router /api/v1/settings [get]
func GetSettings(c *gin.Context) {
	// Retrieve all settings from MongoDB
	settingService := services.SettingService{
		AppConfig: config.Config,
	}
	settings, err := settingService.GetSettings(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	c.JSON(http.StatusOK, dtos.Response{Data: settings})
}
