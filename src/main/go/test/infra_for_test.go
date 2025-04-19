package test

import (
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/root"
	"testing"
)

func TestMain(m *testing.M) {

	if infra.ConfigLogger() {
		return
	}

	var ginServerConfig = infra.GinServerConfiguration{
		Port:    "8081",
		ApiOnly: true,
	}
	var testConfig = infra.Configuration{
		GinServerConfig: ginServerConfig,
	}

	var app = root.App{}
	app.StartOverridingConfigs(&testConfig)
}
