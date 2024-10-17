package infra

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"strings"
	"syscall"
	"time"
)

type GinServer struct {
	router     *gin.Engine
	config     Configuration
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

	glog.Infof("Serving static source files at %s from %s", server.config.webStaticSourceRelPath, server.config.webStaticSourcePath)
	server.router.Static(server.config.webStaticSourceRelPath, server.config.webStaticSourcePath)

	var rootHTMLPath = server.config.webStaticContentPath + "/" + server.config.rootHTMLFilename
	glog.Infof("Serving root HTML file at / from  %s", rootHTMLPath)
	server.router.StaticFile("/", rootHTMLPath)

	glog.Infof("Serving .js, .js.map, .css files from root to load bundles")
	server.router.GET("/:filepath", func(context *gin.Context) {
		file := context.Param("filepath")
		if strings.HasSuffix(file, ".js") ||
			strings.HasSuffix(file, ".js.map") ||
			strings.HasSuffix(file, ".css") {
			context.File(filepath.Join(server.config.webStaticContentPath, file))
		} else {
			handleGetAsToRoot(context, server)
		}
	})

	glog.Infof("All unexpected requests will return the root HTML for front-end routing")
	server.router.NoRoute(func(context *gin.Context) {
		if context.Request.Method == "GET" {
			handleGetAsToRoot(context, server)
		} else {
			context.String(404, "Not Found")
		}
	})
}

func handleGetAsToRoot(context *gin.Context, srvr *GinServer) {
	context.Request.URL.Path = "/"
	srvr.router.HandleContext(context)
}

func (server *GinServer) configRESTRoutes(routes []RESTRoute) {
	for _, route := range routes {
		server.router.Handle(route.Method, route.Path, route.Handlers...)
	}
}

func (server *GinServer) configControllerRoutes(controllers []GinServerRESTController) {
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

func (server *GinServer) configAndWaitForStopSignal() {

	var stopChannel = make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM)

	<-stopChannel
	glog.Info("STOPPING server...")

	contextInstance, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.httpServer.Shutdown(contextInstance); err != nil {
		glog.Fatal("Server forced to shutdown:", err)
	}

	glog.Info("Exiting application process")
}

func (server *GinServer) Init(controllers []GinServerRESTController) {

	glog.Infof("Configuring server before initialization")

	glog.Infof("Configuring basic routes")
	server.configBasicRoutes()

	glog.Infof("Configuring controller routes")
	server.configControllerRoutes(controllers)

	glog.Infof("STARTING server")
	server.start()

	server.configAndWaitForStopSignal()
}

func BuildGinServer(config Configuration) *GinServer {
	return &GinServer{
		router: gin.Default(),
		config: config,
	}
}
