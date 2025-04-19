package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetAllocationPlans(t *testing.T) { //TODO complete infra and test

	//TODO setup test data
	response, err := http.Get("http://localhost:8081/api/portfolio/1/allocation-plan")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	// TODO continue
}
