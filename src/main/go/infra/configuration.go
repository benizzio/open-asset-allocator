package infra

import "os"

type GinServerConfiguration struct {
	port                   string
	webStaticContentPath   string
	webStaticSourceRelPath string
	rootHTMLFilename       string
	webStaticSourcePath    string
	apiRootPath            string
}

type RDBMSConfiguration struct {
	driverName string
	rdbmsURL   string
}

type Configuration struct {
	ginServerConfig GinServerConfiguration
	rdbmsConfig     RDBMSConfiguration
}

func ReadConfig() Configuration {

	var tempWebStaticContentPath = os.Getenv("WEB_STATIC_CONTENT_PATH")
	var tempWebStaticSourceRelPath = os.Getenv("WEB_STATIC_SOURCE_REL_PATH")

	return Configuration{
		ginServerConfig: GinServerConfiguration{
			port:                   os.Getenv("PORT"),
			webStaticContentPath:   tempWebStaticContentPath,
			webStaticSourceRelPath: tempWebStaticSourceRelPath,
			rootHTMLFilename:       os.Getenv("ROOT_HTML_FILENAME"),
			webStaticSourcePath:    tempWebStaticContentPath + tempWebStaticSourceRelPath,
			apiRootPath:            "/api",
		},
		rdbmsConfig: RDBMSConfiguration{
			driverName: os.Getenv("RDBMS_DRIVER_NAME"),
			rdbmsURL:   os.Getenv("RDBMS_URL"),
		},
	}
}
