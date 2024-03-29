// health_check_test.go

package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestReadinessHealthCheckHandler(t *testing.T) {
	// Define a response recorder to capture the HTTP response.
	w := httptest.NewRecorder()

	// Create a GET request to the "/health" endpoint.
	req, err := http.NewRequest("GET", "/readiness", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a Gin context with the request and response recorder.
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Call the healthCheckHandler function from the server package.
	HealthCheckHandler(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"status": "OK"}`, w.Body.String())
}
