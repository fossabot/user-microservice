package api

import (
	"net/http"
	"time"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/thomaspoignant/user-microservice/entity"
)

type UserController struct{}

var userService *entity.UserService

// getUserService init the object to UserService
func userServiceDao() *entity.UserService {
	if userService == nil {
		service, err := entity.NewUserService(viper.GetString("DYNAMODB_TABLE_NAME"))
		if err != nil {
			panic("Impossible to init User service : " + err.Error())
		}
		userService = service
	}
	return userService
}

// Status return a User Object
// @Summary test swagger
// @Description
// @Tags user
// @Consume json
// @Produce  json
// @Param id path int true "Bottle ID"
// @Success 200 {object} entity.User
// @Error 400
// @Router /v1/user/{id} [get]
func (h UserController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, entity.User{
		ID:        uuid.NewV4().String(),
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

// Create a new User
// @Summary test swagger 2
// @Description
// @Tags user
// @Param user body entity.User true "a complete user json"
// @Consume json
// @Produce  json
// @Success 201 {object} entity.User
// @Failure 500 {object} api.ApiErrorResponse "System error"
// @Router /v1/user/ [post]
func (h UserController) Create(c *gin.Context) {
	userToInsert := entity.User{
		FirstName: "John",
		LastName:  "Doe",
	}
	err := userServiceDao().Save(&userToInsert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, userToInsert)
}
