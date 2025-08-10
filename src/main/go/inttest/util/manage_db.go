package util

import (
	"fmt"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	"github.com/golang/glog"
)

type testFormattableSQLPair struct {
	formattableSQL string
	params         []any
}

type CleanupFunctionBuilder struct {
	cleanupQueries []*testFormattableSQLPair
}

func (builder *CleanupFunctionBuilder) AddCleanupQuery(formattableSQL string, params ...any) *CleanupFunctionBuilder {
	builder.cleanupQueries = append(
		builder.cleanupQueries, &testFormattableSQLPair{
			formattableSQL: formattableSQL,
			params:         params,
		},
	)
	return builder
}

func (builder *CleanupFunctionBuilder) Build() func() {
	return createDBCleanupFunctionMulti(builder.cleanupQueries)
}

func BuildCleanupFunctionBuilder() *CleanupFunctionBuilder {
	return &CleanupFunctionBuilder{
		cleanupQueries: make([]*testFormattableSQLPair, 0),
	}
}

func CreateDBCleanupFunction(formattableSQL string, params ...any) func() {
	return createDBCleanupFunctionMulti([]*testFormattableSQLPair{{formattableSQL, params}})
}

func createDBCleanupFunctionMulti(cleanupQueries []*testFormattableSQLPair) func() {
	return func() {
		for _, query := range cleanupQueries {
			glog.Infof("Executing test cleanup query: %s", query.formattableSQL)
			err := inttestinfra.ExecuteDBQuery(fmt.Sprintf(query.formattableSQL, query.params...))
			if err != nil {
				glog.Errorf("Error executing cleanup query: %s", err)
			}
		}
	}
}
