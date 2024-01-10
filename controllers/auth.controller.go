// controllers/auth_controller.go
package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/chokey2nv/gigmile/config"
	"github.com/chokey2nv/gigmile/dtos"
	"github.com/chokey2nv/gigmile/models"
	"github.com/chokey2nv/gigmile/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	ErrToken = errors.New("unable to generate access token")
)

// GenerateToken generates a new JWT token for the given user
func GenerateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires in 1 hour

	return token.SignedString(config.Config.SecretKey)
}

// ParseToken parses the JWT token from the Authorization header
func ParseToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return nil, jwt.ErrSignatureInvalid
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.Config.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

// ParseTokenWithUser parses the JWT token from the Authorization header
// and extracts the user information, storing it in the request context
func ParseTokenWithUser(c *gin.Context) (*jwt.Token, *models.User, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return nil, nil, jwt.ErrSignatureInvalid
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.Config.SecretKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	// Extract user information from token claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, jwt.ErrSignatureInvalid
	}

	// Extract user from claims
	userClaim, ok := claims["user"].(map[string]interface{})
	if !ok {
		return nil, nil, jwt.ErrSignatureInvalid
	}
	var user models.User
	config.ToJSONStruct(userClaim, &user)
	// Store user in the request context
	c.Set("userId", user.Id)

	return token, &user, nil
}

// AuthMiddleware is a middleware to authenticate requests using JWT tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip authentication for signup and login routes
		if c.FullPath() == "/swagger/*any" || c.FullPath() == "/api/v1/signup" || c.FullPath() == "/api/v1/login" {
			c.Next()
			return
		}
		token, _, err := ParseTokenWithUser(c)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// @Summary Signup a new user
// @Description Signup a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dtos.SignUpDto true "User information"
// @Success 200 {object} models.User "User registered successfully"
// @Failure 400 {object} dtos.ErrorResponse "Bad Request"
// @Router /signup [post]
func Signup(c *gin.Context) {
	var signUpDto dtos.SignUpDto
	if err := c.BindJSON(&signUpDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userService := services.UserService{
		AppConfig: config.Config,
	}
	user, err := userService.CreateUser(c, &signUpDto.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dtos.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	userResponseDto := dtos.UpdateUserResponseDto{}
	config.ToJSONStruct(user, &userResponseDto)

	token, err := GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dtos.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: ErrToken.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dtos.Response{Data: gin.H{"user": userResponseDto, "access_token": token}})
}

// @Summary Login with username and password
// @Description Login with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dtos.LoginDto true "User credentials"
// @Success 200 {string} string "Login successful"
// @Failure 401 {object} dtos.ErrorResponse "Unauthorized"
// @Router /login [post]
func Login(c *gin.Context) {
	var loginDto dtos.LoginDto
	if err := c.BindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userService := services.UserService{
		AppConfig: config.Config,
	}
	user, err := userService.GetUser(c, dtos.GetUserDto{Email: loginDto.Email})
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}
	token, err := GenerateToken(user)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			dtos.ErrorResponse{Code: http.StatusInternalServerError, Message: err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
