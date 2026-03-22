package util

import (
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang/glog"
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

func (builder *CleanupFunctionBuilder) Build() func() {
	return createDBCleanupFunctionMulti(builder.cleanupQueries)
}

func BuildCleanupFunctionBuilder() *CleanupFunctionBuilder {
	return &CleanupFunctionBuilder{
		cleanupQueries: make([]*testSQLParamsPair, 0),
	}
}

// CreateDBCleanupFunction creates a cleanup function for a single parameterized SQL query.
// Uses ozzo-dbx named parameter binding ({:paramName} placeholders).
//
// Co-authored by: GitHub Copilot
func CreateDBCleanupFunction(sql string, params dbx.Params) func() {
	return createDBCleanupFunctionMulti([]*testSQLParamsPair{{sql, params}})
}

func createDBCleanupFunctionMulti(cleanupQueries []*testSQLParamsPair) func() {
	return func() {
		for _, query := range cleanupQueries {
			glog.Infof("Executing test cleanup query: %s", query.sql)
			err := inttestinfra.ExecuteDBQuery(query.sql, query.params)
			if err != nil {
				glog.Errorf("Error executing cleanup query: %s", err)
			}
		}
	}
}
