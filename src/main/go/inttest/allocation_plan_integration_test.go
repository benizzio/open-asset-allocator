package inttest

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGetAllocationPlans(t *testing.T) { //TODO complete infra and test

	response, err := http.Get("http://localhost:8081/api/portfolio/1/allocation-plan")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	// TODO continue with assertions
	fmt.Println(string(body))
}
