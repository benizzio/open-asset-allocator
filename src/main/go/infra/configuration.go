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
			ApiOnly:                false,
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
