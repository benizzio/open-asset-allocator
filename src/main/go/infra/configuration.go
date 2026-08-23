package infra

import (
	"os"

	"github.com/benizzio/open-asset-allocator/langext"
)

const defaultYahooFinanceSearchURL = "https://query2.finance.yahoo.com/v1/finance/search"
const defaultYahooFinanceChartURL = "https://query2.finance.yahoo.com/v8/finance/chart/"

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
	DriverName string
	RdbmsURL   string
}

type YahooFinanceConfiguration struct {
	SearchURL string
	ChartURL  string
}

type IntegrationConfiguration struct {
	YahooFinanceConfig YahooFinanceConfiguration
}

type Configuration struct {
	GinServerConfig   GinServerConfiguration
	RdbmsConfig       RDBMSConfiguration
	IntegrationConfig IntegrationConfiguration
}

func (config *Configuration) String() string {
	return langext.StructString(config)
}

// ReadConfig reads the supported process environment settings into a new application
// configuration each time it is called. Callers must set environment variables before
// invoking ReadConfig.
//
// The supported environment settings and defaults are:
//   - PORT defaults to an empty string.
//   - WEB_STATIC_CONTENT_PATH defaults to an empty string.
//   - WEB_STATIC_SOURCE_REL_PATH defaults to an empty string.
//   - ROOT_HTML_FILENAME defaults to an empty string.
//   - RDBMS_DRIVER_NAME defaults to an empty string.
//   - RDBMS_URL defaults to an empty string.
//   - API_ONLY defaults to false and is true only when its value is exactly the
//     lowercase, case-sensitive string "true"; every other value is false.
//   - YAHOO_FINANCE_SEARCH_URL defaults to
//     https://query2.finance.yahoo.com/v1/finance/search when empty.
//   - YAHOO_FINANCE_CHART_URL defaults to
//     https://query2.finance.yahoo.com/v8/finance/chart/ when empty.
//
// The web static source path is derived by concatenating WEB_STATIC_CONTENT_PATH and
// WEB_STATIC_SOURCE_REL_PATH. The API root path is fixed at /api and is not configurable.
// For example, with PORT=8080 and API_ONLY=true set in the process environment before
// the call:
//
//	config := ReadConfig()
//	fmt.Printf("%s %t\n", config.GinServerConfig.Port, config.GinServerConfig.ApiOnly)
//	// Output: 8080 true
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
func ReadConfig() *Configuration {

	var tempWebStaticContentPath = os.Getenv("WEB_STATIC_CONTENT_PATH")
	var tempWebStaticSourceRelPath = os.Getenv("WEB_STATIC_SOURCE_REL_PATH")

	var yahooFinanceSearchURL = os.Getenv("YAHOO_FINANCE_SEARCH_URL")
	if yahooFinanceSearchURL == "" {
		yahooFinanceSearchURL = defaultYahooFinanceSearchURL
	}

	var yahooFinanceChartURL = os.Getenv("YAHOO_FINANCE_CHART_URL")
	if yahooFinanceChartURL == "" {
		yahooFinanceChartURL = defaultYahooFinanceChartURL
	}

	return &Configuration{
		GinServerConfig: GinServerConfiguration{
			Port:                   os.Getenv("PORT"),
			webStaticContentPath:   tempWebStaticContentPath,
			webStaticSourceRelPath: tempWebStaticSourceRelPath,
			rootHTMLFilename:       os.Getenv("ROOT_HTML_FILENAME"),
			webStaticSourcePath:    tempWebStaticContentPath + tempWebStaticSourceRelPath,
			apiRootPath:            "/api",
			ApiOnly:                os.Getenv("API_ONLY") == "true",
		},
		RdbmsConfig: RDBMSConfiguration{
			DriverName: os.Getenv("RDBMS_DRIVER_NAME"),
			RdbmsURL:   os.Getenv("RDBMS_URL"),
		},
		IntegrationConfig: IntegrationConfiguration{
			YahooFinanceConfig: YahooFinanceConfiguration{
				SearchURL: yahooFinanceSearchURL,
				ChartURL:  yahooFinanceChartURL,
			},
		},
	}
}
