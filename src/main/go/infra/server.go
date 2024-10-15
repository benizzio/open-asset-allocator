package infra

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"path/filepath"
	"strings"
)

type GinServer struct {
	router *gin.Engine
	config Configuration
}

func (server *GinServer) configBasicRoutes() {

	glog.Info("Configuring Gin router")

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

func (server *GinServer) start() {
	err := server.router.Run(fmt.Sprintf(":%s", server.config.port))
	if err != nil {
		glog.Error("Error starting server: ", err)
	}
}

func (server *GinServer) Init() {
	server.configBasicRoutes()
	server.start()
}

func BuildGinServer(config Configuration) *GinServer {
	return &GinServer{
		router: gin.Default(),
		config: config,
	}
}
