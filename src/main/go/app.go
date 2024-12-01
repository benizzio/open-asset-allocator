package main

import (
	"context"
	"errors"
	"github.com/benizzio/open-asset-allocator/api/rest"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/domain/infra/repository"
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

	//Change this to a fluent API
	databaseAdapter, server := initializeAppComponents(buildAndInjectAppComponents())

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

func buildAndInjectAppComponents() (*infra.RDBMSAdapter, *infra.GinServer, []infra.GinServerRESTController) {

	var config = infra.ReadConfig()
	var server = infra.BuildGinServer(config)
	var databaseAdapter = infra.BuildDatabaseAdapter(config)

	var portfolioRepository = repository.BuildPortfolioRepository(databaseAdapter)
	var portfolioHistoryService = application.BuildPortfolioHistoryService(portfolioRepository)
	var portfolioRESTController = rest.BuildPortfolioRESTController(portfolioHistoryService)

	var allocationPlanRepository = repository.BuildAllocationPlanRepository(databaseAdapter)
	var allocationPlanService = application.BuildAllocationPlanService(allocationPlanRepository)
	var allocationRESTController = rest.BuildAllocationRESTController(allocationPlanService)

	var restControllers = []infra.GinServerRESTController{portfolioRESTController, allocationRESTController}

	return databaseAdapter, server, restControllers
}

func initializeAppComponents(
	databaseAdapter *infra.RDBMSAdapter,
	server *infra.GinServer,
	restControllers []infra.GinServerRESTController,
) (
	*infra.RDBMSAdapter,
	*infra.GinServer,
) {

	databaseAdapter.Init()
	databaseAdapter.Ping()

	server.Init(restControllers)

	return databaseAdapter, server
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
