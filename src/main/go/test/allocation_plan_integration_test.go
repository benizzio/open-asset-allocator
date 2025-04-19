package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetAllocationPlans(t *testing.T) { //TODO complete infra and test

	router := SetUpRouter()

	request, err := http.NewRequest(http.MethodGet, "http://localhost:8081/api/portfolio/1/allocation-plan", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}
