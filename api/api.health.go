package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController who handle health check
type HealthController struct {
	Status int    `json:"status" example:"200"`
	Code   string `json:"code" example:"SUCCESS"`
	Health string `json:"health" example:"RUNNING"`
}

// HealthCheck return the Status of the current app
// @Summary Health check endpoint
// @Description health check endpoint to know if the service is up
// @Tags healthcheck
// @Produce  json
// @Success 200 {object} api.HealthController
// @Router /health [get]
func (h HealthController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthController{
		Status: http.StatusOK,
		Code:   Success,
		Health: "RUNNING",
	})
}
