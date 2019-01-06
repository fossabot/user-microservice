package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomaspoignant/user-microservice/testingUtils"
)

// TestHealthCheck check if the health check is answering correctly
func Test404(t *testing.T) {
	testingUtils.PrepareTest()
	router := SetupRouter()

	w := testingUtils.PerformHTTPRequest(router, "GET", "/not_exiting_path")
	assert.Equal(t, http.StatusNotFound, w.Code)

	expected := ApiErrorResponse{
		Status:  http.StatusNotFound,
		Code:    NotFound,
		Message: "Resource not found",
	}

	var got ApiErrorResponse
	err := json.Unmarshal([]byte(w.Body.String()), &got)
	assert.Nil(t, err)
	assert.Equal(t, expected, got)
}
