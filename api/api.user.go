package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/thomaspoignant/user-microservice/entity"
)

type UserController struct{}

// User is the response object of the API
type user struct {
	// API return code
	Code string `json:"code" example:"SUCCESS"`
	// informations of a user
	User entity.User `json:"user"`
}

// Status return a User Object
// @Summary test swagger
// @Description
// @Tags user
// @Consume json
// @Produce  json
// @Param id path int true "Bottle ID"
// @Success 200 {object} api.user
// @Error 400
// @Router /v1/user/{id} [get]
func (h UserController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, user{
		Code: "SUCCESS",
		User: entity.User{
			ID:        uuid.NewV4().String(),
			FirstName: "John",
			LastName:  "Doe",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
}

func (h UserController) Status2(c *gin.Context) {
	c.String(http.StatusOK, "Working!2")
}
