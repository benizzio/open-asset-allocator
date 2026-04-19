package rdbms

import (
	"database/sql"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang/glog"
)

// ================================================
// QUERY BUILDER
// ================================================

// QueryBuilder builds parameterized SQL queries with optional WHERE clause composition.
// The type parameter T propagates to the resulting QueryExecutor, enabling type-safe
// result scanning.
//
// Example:
//
//	assets, err := rdbms.BuildQuery[domain.Asset](adapter, assetsSQL).
//		AddWhereClauseAndParam("AND id = {:id}", "id", 42).
//		Build().
//		FindWithRowScanner(assetRowScanner)
//
// Co-authored by: GitHub Copilot and Igor Benicio de Mesquita
type QueryBuilder[T any] struct {
	dbx          *dbx.DB
	querySQL     string
	whereClauses []string
	params       dbx.Params
}

// Build finalizes the query builder and returns a QueryExecutor ready for execution.
//
// Co-authored by: GitHub Copilot and Igor Benicio de Mesquita
func (builder *QueryBuilder[T]) Build() *QueryExecutor[T] {

	var processedSQL = processSQL(builder.querySQL, builder.whereClauses)

	var query = builder.dbx.NewQuery(processedSQL)
	var queryExecutor = withParams[T](query, builder.params)
	return queryExecutor
}

func (builder *QueryBuilder[T]) AddParam(name string, value any) *QueryBuilder[T] {
	builder.params[name] = value
	return builder
}

func (builder *QueryBuilder[T]) AddWhereClause(whereClause string) *QueryBuilder[T] {
	builder.whereClauses = append(builder.whereClauses, whereClause)
	return builder
}

func (builder *QueryBuilder[T]) AddWhereClauseAndParam(
	whereClause string,
	name string,
	value any,
) *QueryBuilder[T] {
	return builder.AddWhereClause(whereClause).AddParam(name, value)
}

// ================================================
// QUERY EXECUTOR
// ================================================

// QueryExecutor executes parameterized queries built through QueryBuilder, providing type-safe
// result scanning via generics. Supports both ozzo-dbx struct mapping (FindInto, GetInto) and
// custom row scanners (FindWithRowScanner, GetWithRowScanner).
//
// Example:
//
//	executor := rdbms.BuildQuery[domain.Asset](adapter, assetsSQL).Build()
//	assets, err := executor.FindWithRowScanner(assetRowScanner)
//
// Co-authored by: GitHub Copilot and Igor Benicio de Mesquita
type QueryExecutor[T any] struct {
	query *dbx.Query
}

func withParams[T any](query *dbx.Query, params dbx.Params) *QueryExecutor[T] {
	if len(params) > 0 {
		query.Bind(params)
	}
	return &QueryExecutor[T]{query: query}
}

func (executor *QueryExecutor[T]) FindInto(target *[]T) error {
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Executing query %s \n with params %v", executor.query.SQL(), executor.query.Params())
	return executor.query.All(target)
}

func (executor *QueryExecutor[T]) GetInto(target *T) error {
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Executing query %s \n with params %v", executor.query.SQL(), executor.query.Params())
	return executor.query.One(target)
}

func (executor *QueryExecutor[T]) GetRows() (*dbx.Rows, error) {
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Executing query %s \n with params %v", executor.query.SQL(), executor.query.Params())
	return executor.query.Rows()
}

// FindWithRowScanner executes the query and maps each result row with the provided scanner.
//
// Example:
//
//	queryExecutor := rdbms.BuildQuery[domain.Asset](
//		adapter, "SELECT id, ticker FROM asset",
//	).Build()
//	assets, err := queryExecutor.FindWithRowScanner(assetRowScanner)
//
// Co-authored by: GitHub Copilot, OpenCode and Igor Benicio de Mesquita
func (executor *QueryExecutor[T]) FindWithRowScanner(rowScanner RowScanner[T]) ([]T, error) {
	rows, err := executor.GetRows()
	if err != nil {
		return nil, err
	}

	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			glog.Errorf("Error closing rows: %v", closeErr)
		}
	}()

	var result = make([]T, 0)
	var index = 0
	for rows.Next() {
		rowValue, scanErr := rowScanner(rows.Rows)
		if scanErr != nil {
			glog.Errorf("Error scanning row %d: %v", index, scanErr)
			return nil, scanErr
		}

		result = append(result, rowValue)
		index++
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// GetWithRowScanner executes the query and maps the first result row with the provided
// scanner. When the query returns no rows, sql.ErrNoRows is returned.
//
// Example:
//
//	queryExecutor := rdbms.BuildQuery[domain.Asset](
//		adapter, "SELECT id, ticker FROM asset WHERE id = {:id}",
//	).
//		AddParam("id", 1).
//		Build()
//	asset, err := queryExecutor.GetWithRowScanner(assetRowScanner)
//
// Co-authored by: GitHub Copilot, OpenCode and Igor Benicio de Mesquita
func (executor *QueryExecutor[T]) GetWithRowScanner(rowScanner RowScanner[T]) (T, error) {
	var zero T

	rows, err := executor.GetRows()
	if err != nil {
		return zero, err
	}

	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			glog.Errorf("Error closing rows: %v", closeErr)
		}
	}()

	if !rows.Next() {
		if err = rows.Err(); err != nil {
			return zero, err
		}
		return zero, sql.ErrNoRows
	}

	return rowScanner(rows.Rows)
}
