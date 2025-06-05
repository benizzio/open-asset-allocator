package inttest

import (
	"context"
	"github.com/benizzio/open-asset-allocator/infra/util"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// TODO tests are running successfully when migrations fail, FIX
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
