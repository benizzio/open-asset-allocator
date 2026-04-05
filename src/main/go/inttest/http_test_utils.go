package inttest

import (
	"net/http"

	"github.com/golang/glog"
)

func deferCloseResponseBody(response *http.Response) {
	err := response.Body.Close()
	if err != nil {
		glog.Errorf("Error closing response body: %v", err)
	}
}
