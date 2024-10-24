package main

import (
	"context"
	"errors"
	"github.com/benizzio/open-asset-allocator/api/rest"
	"github.com/benizzio/open-asset-allocator/application"
	domain_infra "github.com/benizzio/open-asset-allocator/domain/infra/repository"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TODO clean
func main() {

	if infra.ConfigLogger() {
		return
	}

	server, databaseAdapter := buildAppComponents()

	var portfolioRepository = domain_infra.BuildPortfolioRepository(databaseAdapter)
	var portfolioHistoryService = application.BuildPortfolioHistoryService(portfolioRepository)
	var portfolioRESTController = rest.BuildPortfolioRESTController(portfolioHistoryService)

	initializeAppComponents(databaseAdapter, server, []infra.GinServerRESTController{portfolioRESTController})

	databaseAdapter.Ping()

	stopChannel := buildStopChannel()

	<-stopChannel

	stopContext, cancel := buildStopContext()
	defer cancel()

	closeAppComponents(stopContext, server, databaseAdapter)

	glog.Info("Exiting application process")
}

func buildStopContext() (context.Context, context.CancelFunc) {
	stopContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	go func() {
		<-stopContext.Done()
		if errors.Is(stopContext.Err(), context.DeadlineExceeded) {
			glog.Fatal("Error stopping application: timeout")
		}
	}()
	return stopContext, cancel
}

func buildAppComponents() (*infra.GinServer, *infra.RDBMSAdapter) {
	var config = infra.ReadConfig()
	var server = infra.BuildGinServer(config)
	var databaseAdapter = infra.BuildDatabaseAdapter(config)
	return server, databaseAdapter
}

func initializeAppComponents(
	databaseAdapter *infra.RDBMSAdapter,
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
	stopContext context.Context,
	server *infra.GinServer,
	databaseAdapter *infra.RDBMSAdapter,
) {
	server.Stop(stopContext)
	databaseAdapter.Stop()
}
