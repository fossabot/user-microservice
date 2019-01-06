package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (h UserController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!%s")

}

func (h UserController) Status2(c *gin.Context) {
	c.String(http.StatusOK, "Working!2")
}
