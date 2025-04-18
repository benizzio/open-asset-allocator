package infra

import "os"

type GinServerConfiguration struct {
	port                   string
	webStaticContentPath   string
	webStaticSourceRelPath string
	rootHTMLFilename       string
	webStaticSourcePath    string
	apiRootPath            string
	apiOnly                bool
}

type RDBMSConfiguration struct {
	driverName string
	rdbmsURL   string
}

type Configuration struct {
	GinServerConfig GinServerConfiguration
	RdbmsConfig     RDBMSConfiguration
}

func ReadConfig() *Configuration {

	var tempWebStaticContentPath = os.Getenv("WEB_STATIC_CONTENT_PATH")
	var tempWebStaticSourceRelPath = os.Getenv("WEB_STATIC_SOURCE_REL_PATH")

	return &Configuration{
		GinServerConfig: GinServerConfiguration{
			port:                   os.Getenv("PORT"),
			webStaticContentPath:   tempWebStaticContentPath,
			webStaticSourceRelPath: tempWebStaticSourceRelPath,
			rootHTMLFilename:       os.Getenv("ROOT_HTML_FILENAME"),
			webStaticSourcePath:    tempWebStaticContentPath + tempWebStaticSourceRelPath,
			apiRootPath:            "/api",
			apiOnly:                false,
		},
		RdbmsConfig: RDBMSConfiguration{
			driverName: os.Getenv("RDBMS_DRIVER_NAME"),
			rdbmsURL:   os.Getenv("RDBMS_URL"),
		},
	}
}
