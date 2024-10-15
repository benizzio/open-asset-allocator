package main

import (
	"github.com/benizzio/open-asset-allocator/infra"
)

func main() {

	if infra.ConfigLogger() {
		return
	}

	var config = infra.ReadConfig()
	var srvr = infra.BuildGinServer(config)
	srvr.Init()
}
