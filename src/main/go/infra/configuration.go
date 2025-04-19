package infra

import "os"

type GinServerConfiguration struct {
	Port                   string
	webStaticContentPath   string
	webStaticSourceRelPath string
	rootHTMLFilename       string
	webStaticSourcePath    string
	apiRootPath            string
	ApiOnly                bool
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
			Port:                   os.Getenv("PORT"),
			webStaticContentPath:   tempWebStaticContentPath,
			webStaticSourceRelPath: tempWebStaticSourceRelPath,
			rootHTMLFilename:       os.Getenv("ROOT_HTML_FILENAME"),
			webStaticSourcePath:    tempWebStaticContentPath + tempWebStaticSourceRelPath,
			apiRootPath:            "/api",
			ApiOnly:                false,
		},
		RdbmsConfig: RDBMSConfiguration{
			driverName: os.Getenv("RDBMS_DRIVER_NAME"),
			rdbmsURL:   os.Getenv("RDBMS_URL"),
		},
	}
}
