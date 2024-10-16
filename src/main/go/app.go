package main

import (
	"github.com/benizzio/open-asset-allocator/api"
	"github.com/benizzio/open-asset-allocator/infra"
)

func main() {

	if infra.ConfigLogger() {
		return
	}

	var config = infra.ReadConfig()
	var server = infra.BuildGinServer(config)

	var portfolioRESTController = api.BuildPortfolioRESTController()

	server.Init([]infra.GinServerRESTController{portfolioRESTController})
}
