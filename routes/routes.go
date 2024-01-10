// routes/routes.go
package routes

import (
	_ "github.com/chokey2nv/gigmile/docs"

	"github.com/chokey2nv/gigmile/config"
	"github.com/chokey2nv/gigmile/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router variable to hold the gin.Engine instance (single)
var Router *gin.Engine

// Run initializes and runs the API
func Run() {
	// Load configuration
	config.LoadConfig()

	// Initialize Gin router
	Router = gin.Default()

	// Apply middleware and routes
	Router.Use(cors.Default())

	// Apply middleware and routes
	Router.Use(controllers.AuthMiddleware())

	/*
		for extra security, you can put this server behind a proxy
		Explicitly set trusted proxy headers
		r.ForwardedByClientIP = true
		r.SetTrustedProxies(proxyIP || ipRange)
	*/

	v1 := Router.Group("/api/v1")
	updates := v1.Group("/updates")
	{
		// Post update
		updates.POST("/new", controllers.CreateUpdate)
		// Get updates
		updates.POST("/get", controllers.GetUpdates)
	}
	sprints := v1.Group("/sprints")
	{
		sprints.POST("/new", controllers.CreateSprint)
		sprints.POST("/get", controllers.GetSprints)
	}
	tasks := v1.Group("/tasks")
	{
		tasks.POST("/new", controllers.CreateTask)
		tasks.POST("/get", controllers.GetTasks)
	}
	users := v1.Group("/users")
	{
		users.GET("/", controllers.GetUsers)
	}
	settings := v1.Group("/settings")
	{
		settings.POST("/", controllers.CreateSetting)
		settings.GET("/", controllers.GetSettings)
	}
	// Routes for authentication
	v1.POST("/signup", controllers.Signup)
	v1.POST("/login", controllers.Login)

	// Swagger documentation route
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server
	if err := Router.Run(":8080"); err != nil {
		panic(err)
	}
}
