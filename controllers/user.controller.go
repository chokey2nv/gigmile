package controllers

import (
	"net/http"

	"github.com/chokey2nv/gigmile/config"
	"github.com/chokey2nv/gigmile/dtos"
	"github.com/chokey2nv/gigmile/services"
	"github.com/gin-gonic/gin"
)

// GetUsers handles retrieving all users
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Param input body dtos.GetUserDto true "Page information"
// @Success 200 {array} []models.User "List of users"
// @Failure 500 {object} dtos.ErrorResponse "Internal Server Error"
// @Router /api/v1/users [get]
func GetUsers(c *gin.Context) {
	var getUsersDto dtos.GetUsersDto
	if err := c.BindJSON(&getUsersDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userService := &services.UserService{
		AppConfig: config.Config,
	}
	users, err := userService.GetUsers(c, getUsersDto.PageOption)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dtos.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &dtos.Response{Data: users})
}
