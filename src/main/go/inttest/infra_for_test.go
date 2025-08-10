package inttest

import (
	"context"
	"os"
	"testing"

	"github.com/benizzio/open-asset-allocator/infra/util"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
)

func TestMain(m *testing.M) {

	ctx := context.Background()

	var deferRegistry = util.BuildDeferRegistry()
	ctx = context.WithValue(ctx, inttestinfra.DeferRegistryKey, deferRegistry)

	var exitVal = setupAndRunTests(ctx, m)

	deferRegistry.Execute()
	os.Exit(exitVal)
}

func setupAndRunTests(ctx context.Context, m *testing.M) int {

	var exitVal = inttestinfra.SetupTestInfra(ctx)
	if exitVal != 0 {
		return exitVal
	}

	var app = inttestinfra.BuildAndStartApplication()
	defer func() {
		app.Stop()
	}()

	// run tests
	return m.Run()
}
