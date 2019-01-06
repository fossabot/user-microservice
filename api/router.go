package api

import (
	"net/http"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

// SetupRouter determine all the routes for this service
func SetupRouter() *gin.Engine {
	// setting Gin mode before running
	gin.SetMode(viper.GetString("GIN_MODE"))

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//router.Use(middlewares.AuthMiddleware())

	// healthCheck router
	health := new(HealthController)
	router.GET("/health", health.HealthCheck)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(UserController)
			//TODO : mettre les bonnes méthodes en face
			userGroup.GET("/", user.Status2)
			userGroup.GET("/:id", user.Status)
			userGroup.POST("/", user.Status)
			userGroup.PATCH("/:id", user.Status)
			userGroup.DELETE("/:id", user.Status)
		}
	}

	// Returning 404 when calling an unmapped uri
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, ApiErrorResponse{
			Code:    NotFound,
			Status:  http.StatusNotFound,
			Message: "Resource not found",
		})
	})

	return router
}
