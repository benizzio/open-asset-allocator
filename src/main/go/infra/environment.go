package infra

import "os"

type Environment struct {
	port                   string
	webStaticContentPath   string
	webStaticSourceRelPath string
	rootHTMLFilename       string
	webStaticSourcePath    string
}

func ReadEnvironment() Environment {

	var tempWebStaticContentPath = os.Getenv("WEB_STATIC_CONTENT_PATH")
	var tempWebStaticSourceRelPath = os.Getenv("WEB_STATIC_SOURCE_REL_PATH")

	return Environment{
		port:                   os.Getenv("PORT"),
		webStaticContentPath:   tempWebStaticContentPath,
		webStaticSourceRelPath: tempWebStaticSourceRelPath,
		rootHTMLFilename:       os.Getenv("ROOT_HTML_FILENAME"),
		webStaticSourcePath:    tempWebStaticContentPath + tempWebStaticSourceRelPath,
	}
}
