package main

import (
	"github.com/benizzio/open-asset-allocator/infra"
)

func main() {

	if infra.ConfigLogger() {
		return
	}

	var env = infra.ReadEnvironment()
	var srvr = infra.BuildGinServer(env)
	srvr.Init()
}
