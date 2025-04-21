package main

import (
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/root"
)

func main() {

	if infra.ConfigLogger() {
		return
	}

	var app = root.App{}
	app.Run()
}
