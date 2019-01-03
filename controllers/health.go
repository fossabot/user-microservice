package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thomaspoignant/user-microservice/api"
)

// HealthController who handle health check
type HealthController struct {
	Status int    `json:"status"`
	Code   string `json:"code"`
	Health string `json:"health"`
}

// HealthCheck return the Status of the current app
func (h HealthController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthController{
		Status: http.StatusOK,
		Code:   api.Success,
		Health: "RUNNING",
	})
}
