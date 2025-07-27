package util

import (
	"fmt"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	"github.com/golang/glog"
)

func CreateDBCleanupFunction(formattableSQL string, params ...any) func() {
	return func() {
		glog.Infof("Executing test cleanup query: %s", formattableSQL)
		err := inttestinfra.ExecuteDBQuery(fmt.Sprintf(formattableSQL, params...))
		if err != nil {
			glog.Errorf("Error executing cleanup query: %s", err)
		}
	}
}
