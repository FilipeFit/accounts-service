package heath

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request = request
	HealthCheck(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "{\"status\":\"UP\"}", response.Body.String())
}
