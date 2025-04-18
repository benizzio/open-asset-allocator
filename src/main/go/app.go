package main

import (
	"context"
	"errors"
	"github.com/benizzio/open-asset-allocator/api/rest"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/domain/infra/repository"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	config          *infra.Configuration
	databaseAdapter *infra.RDBMSAdapter
	server          *infra.GinServer
	restControllers []infra.GinServerRESTController
}

func (app *App) buildBaseInfrastructure() {
	var config = infra.ReadConfig()
	app.config = config
	app.server = infra.BuildGinServer(config)
	app.databaseAdapter = infra.BuildDatabaseAdapter(config)
}

func (app *App) buildAppComponents() {

	var portfolioRepository = repository.BuildPortfolioRepository(app.databaseAdapter)
	var allocationPlanRepository = repository.BuildAllocationPlanRepository(app.databaseAdapter)

	var portfolioDomService = service.BuildPortfolioDomService(portfolioRepository)
	var allocationPlanDomService = service.BuildAllocationPlanDomService(allocationPlanRepository)
	var portfolioDivergenceAnalysisService = application.BuildPortfolioDivergenceAnalysisAppService(
		portfolioDomService,
		allocationPlanDomService,
	)
	var portfolioAnalysisConfigurationService = application.BuildPortfolioAnalysisConfigurationAppService(
		portfolioDomService,
		allocationPlanDomService,
	)

	var portfolioRESTController = rest.BuildPortfolioRESTController(
		portfolioDomService,
		portfolioAnalysisConfigurationService,
		portfolioDivergenceAnalysisService,
	)
	var allocationPlanRESTController = rest.BuildAllocationPlanRESTController(allocationPlanDomService)

	app.restControllers = []infra.GinServerRESTController{portfolioRESTController, allocationPlanRESTController}
}

func (app *App) initializeAppComponents() {
	app.databaseAdapter.Init()
	app.databaseAdapter.Ping()
	app.server.Init(app.restControllers)
}

func (app *App) closeAppComponents(stopContext context.Context) {
	app.server.Stop(stopContext)
	app.databaseAdapter.Stop()
}

func (app *App) Start() {

	app.buildBaseInfrastructure()
	app.buildAppComponents()
	app.initializeAppComponents()

	stopChannel := buildStopChannel()

	<-stopChannel

	stopContext, cancel := buildStopContext()
	defer cancel()

	app.closeAppComponents(stopContext)

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

func buildStopChannel() chan os.Signal {
	var stopChannel = make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(stopChannel, os.Interrupt, os.Kill)
	return stopChannel
}

func main() {

	if infra.ConfigLogger() {
		return
	}

	var app = App{}
	app.Start()
}
