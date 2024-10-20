package main

import (
	"context"
	"github.com/benizzio/open-asset-allocator/api"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	if infra.ConfigLogger() {
		return
	}

	server, databaseAdapter := buildAppComponents()

	var portfolioRESTController = api.BuildPortfolioRESTController()
	initializeAppComponents(databaseAdapter, server, []infra.GinServerRESTController{portfolioRESTController})

	databaseAdapter.Ping()

	stopChannel := buildStopChannel()

	<-stopChannel

	contextInstance, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	closeAppComponents(contextInstance, server, databaseAdapter, cancel)

	glog.Info("Exiting application process")
}

func buildAppComponents() (*infra.GinServer, *infra.DatabaseAdapter) {
	var config = infra.ReadConfig()
	var server = infra.BuildGinServer(config)
	var databaseAdapter = infra.BuildDatabaseAdapter(config)
	return server, databaseAdapter
}

func initializeAppComponents(
	databaseAdapter *infra.DatabaseAdapter,
	server *infra.GinServer,
	restControllers []infra.GinServerRESTController,
) {
	databaseAdapter.Init()
	server.Init(restControllers)
}

func buildStopChannel() chan os.Signal {
	var stopChannel = make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(stopChannel, os.Interrupt, os.Kill)
	return stopChannel
}

func closeAppComponents(
	contextInstance context.Context,
	server *infra.GinServer,
	databaseAdapter *infra.DatabaseAdapter,
	cancel context.CancelFunc,
) {
	defer cancel()
	server.Stop(contextInstance)
	databaseAdapter.Stop()
}
