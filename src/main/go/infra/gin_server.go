package infra

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
)

type GinServer struct {
	router     *gin.Engine
	config     GinServerConfiguration
	httpServer *http.Server
}

type RESTRoute struct {
	Method   string
	Path     string
	Handlers gin.HandlersChain
}

type GinServerRESTController interface {
	BuildRoutes() []RESTRoute
}

func (server *GinServer) configBasicRoutes() {
	server.configStaticFilesRoute()
	server.configRootHMTLAndDependenciesRoute()
	server.configUnknownStandardRoutes()
}

func (server *GinServer) configStaticFilesRoute() {
	glog.Infof(
		"Serving static source files at %s from %s",
		server.config.webStaticSourceRelPath,
		server.config.webStaticSourcePath,
	)
	server.router.Static(
		server.config.webStaticSourceRelPath,
		server.config.webStaticSourcePath,
	)
}

func (server *GinServer) configRootHMTLAndDependenciesRoute() {

	var rootHTMLPath = server.config.webStaticContentPath + "/" + server.config.rootHTMLFilename
	glog.Infof("Serving root HTML file at / from  %s", rootHTMLPath)
	server.router.StaticFile("/", rootHTMLPath)

	glog.Infof("Serving .js, .js.map, .css files from root to load bundles")
	server.router.GET(
		"/:filepath", func(context *gin.Context) {

			file := context.Param("filepath")

			// Validate the file parameter to prevent path traversal
			if strings.Contains(file, "/") || strings.Contains(file, "\\") || strings.Contains(file, "..") {
				context.String(http.StatusBadRequest, "Invalid file name")
				return
			}

			if strings.HasSuffix(file, ".js") ||
				strings.HasSuffix(file, ".js.map") ||
				strings.HasSuffix(file, ".css") ||
				strings.HasSuffix(file, ".woff") ||
				strings.HasSuffix(file, ".woff2") {
				context.File(filepath.Join(server.config.webStaticContentPath, file))
			} else {
				server.handleGetAsToRoot(context)
			}
		},
	)
}

func (server *GinServer) configUnknownStandardRoutes() {

	glog.Infof("All unexpected requests outside the API path will return the root HTML for front-end routing")
	server.router.NoRoute(
		func(context *gin.Context) {
			if context.Request.Method == "GET" && !server.isRequestToAPI(context) {
				server.handleGetAsToRoot(context)
			} else {
				context.String(http.StatusNotFound, "Not Found")
			}
		},
	)
}

func (server *GinServer) handleGetAsToRoot(context *gin.Context) {
	context.Request.URL.Path = "/"
	server.router.HandleContext(context)
}

func (server *GinServer) isRequestToAPI(context *gin.Context) bool {
	return strings.HasPrefix(context.Request.URL.EscapedPath(), server.config.apiRootPath)
}

func (server *GinServer) configRESTRoutes(routes []RESTRoute) {
	for _, route := range routes {
		server.router.Handle(route.Method, route.Path, route.Handlers...)
	}
}

func (server *GinServer) configControllerRoutes(controllers []GinServerRESTController) {
	glog.Infof("Received %d controllers to config routes", len(controllers))
	for _, controller := range controllers {
		glog.Infof("Configuring controller routes for %s", reflect.TypeOf(controller))
		server.configRESTRoutes(controller.BuildRoutes())
	}
}

func (server *GinServer) start() {

	server.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", server.config.port),
		Handler: server.router,
	}

	go func() {
		if err := server.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			glog.Fatalf("listen: %s\n", err)
		}
	}()
}

func (server *GinServer) stop(stopContext context.Context) {
	if err := server.httpServer.Shutdown(stopContext); err != nil {
		glog.Fatal("Server forced to shutdown: ", err)
	}
}

func (server *GinServer) Init(controllers []GinServerRESTController) {

	glog.Infof("Configuring server before initialization")

	if !server.config.apiOnly {
		glog.Infof("Configuring basic routes")
		server.configBasicRoutes()
	}

	glog.Infof("Configuring controller routes")
	server.configControllerRoutes(controllers)

	glog.Infof("STARTING server ==========>")
	server.start()
}

func (server *GinServer) Stop(stopContext context.Context) {
	glog.Info("STOPPING server <==========")
	server.stop(stopContext)
}

func BuildGinServer(config *Configuration) *GinServer {
	return &GinServer{
		router: gin.Default(),
		config: config.GinServerConfig,
	}
}
