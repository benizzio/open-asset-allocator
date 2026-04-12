package util

import (
	"testing"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang/glog"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
)

type testSQLParamsPair struct {
	sql    string
	params dbx.Params
}

// CleanupFunctionBuilder accumulates parameterized SQL cleanup queries
// and produces a single cleanup function that executes them in order.
//
// Co-authored by: GitHub Copilot
type CleanupFunctionBuilder struct {
	cleanupQueries []*testSQLParamsPair
}

// AddCleanupQuery appends a parameterized SQL cleanup query to the builder.
// Uses ozzo-dbx named parameter binding ({:paramName} placeholders).
//
// Co-authored by: GitHub Copilot
func (builder *CleanupFunctionBuilder) AddCleanupQuery(sql string, params dbx.Params) *CleanupFunctionBuilder {
	builder.cleanupQueries = append(
		builder.cleanupQueries, &testSQLParamsPair{
			sql:    sql,
			params: params,
		},
	)
	return builder
}

// Build produces a cleanup function that executes all accumulated queries in order.
// Cleanup errors are reported via t.Errorf so the test fails but remaining queries still run.
//
// Co-authored by: GitHub Copilot
func (builder *CleanupFunctionBuilder) Build(t *testing.T) func() {
	return createDBCleanupFunctionMulti(t, builder.cleanupQueries)
}

func BuildCleanupFunctionBuilder() *CleanupFunctionBuilder {
	return &CleanupFunctionBuilder{
		cleanupQueries: make([]*testSQLParamsPair, 0),
	}
}

func createDBCleanupFunctionMulti(t *testing.T, cleanupQueries []*testSQLParamsPair) func() {
	return func() {
		for _, query := range cleanupQueries {
			glog.Infof("Executing test cleanup query: %s", query.sql)
			err := inttestinfra.ExecuteDBQuery(query.sql, query.params)
			if err != nil {
				t.Errorf("Error executing cleanup query: %s", err)
			}
		}
	}
}
