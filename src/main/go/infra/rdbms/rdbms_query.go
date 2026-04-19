package rdbms

import (
	"database/sql"
	"strings"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/golang/glog"
	"github.com/lib/pq"

	"github.com/benizzio/open-asset-allocator/langext"
)

const (
	WhereClausePlaceholder = "/*WHERE+PARAMS*/"
)

func processSQL(querySQL string, whereClauses []string) string {

	var processedSQL = querySQL

	if len(whereClauses) > 0 {
		var whereStatement = " WHERE 1=1 " + strings.Join(whereClauses, " ")
		processedSQL = strings.Replace(processedSQL, WhereClausePlaceholder, whereStatement, 1)
	} else {
		processedSQL = strings.Replace(processedSQL, WhereClausePlaceholder, "", 1)
	}

	return processedSQL
}

// processParamsForPostgreSQL converts slice parameters to pq.Array for PostgreSQL compatibility.
//
// Parameters:
//   - params: Variable number of parameters that may include slices
//
// Returns:
//   - []any: Processed parameters with slices converted to pq.Array
//
// Authored by: GitHub Copilot
func processParamsForPostgreSQL(params ...any) []any {

	var processedParams = make([]any, len(params))

	for i, param := range params {
		if langext.IsSlice(param) {
			processedParams[i] = pq.Array(param)
		} else {
			processedParams[i] = param
		}
	}

	return processedParams
}

// ================================================
// QUERY BUILDER
// ================================================

// QueryBuilder builds parameterized SQL queries with optional WHERE clause composition.
// The type parameter T propagates to the resulting QueryExecutor, enabling type-safe
// result scanning.
//
// Example:
//
//	assets, err := rdbms.BuildQuery[domain.Asset](adapter.GetDBX(), assetsSQL).
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

func (builder *QueryBuilder[T]) AddWhereClauseAndParam(whereClause string, name string, value any) *QueryBuilder[T] {
	return builder.AddWhereClause(whereClause).AddParam(name, value)
}

// ================================================
// QUERY BUILDER - sql.Tx
// ================================================

type SQLTransactionalQueryBuilder[T any] struct {
	transaction  *sql.Tx
	querySQL     string
	whereClauses []string
	params       []any
}

func (builder *SQLTransactionalQueryBuilder[T]) AddParams(params ...any) *SQLTransactionalQueryBuilder[T] {
	var processedParams = processParamsForPostgreSQL(params...)
	builder.params = append(builder.params, processedParams...)
	return builder
}

func (builder *SQLTransactionalQueryBuilder[T]) AddWhereClause(whereClause string) *SQLTransactionalQueryBuilder[T] {
	builder.whereClauses = append(builder.whereClauses, whereClause)
	return builder
}

// AddWhereClauseAndParams adds a WHERE clause and its parameters to the query builder.
//
// This method automatically converts slice parameters to pq.Array for PostgreSQL
// compatibility, allowing seamless use of slices in SQL IN clauses and array operations.
//
// Parameters:
//   - whereClause: The WHERE clause SQL fragment to add
//   - params: Variable number of parameters, with slices automatically converted to pq.Array
//
// Returns:
//   - *SQLTransactionalQueryBuilder[T]: The same builder instance for method chaining
//
// Co-authored by: GitHub Copilot
func (builder *SQLTransactionalQueryBuilder[T]) AddWhereClauseAndParams(
	whereClause string,
	params ...any,
) *SQLTransactionalQueryBuilder[T] {

	builder.whereClauses = append(builder.whereClauses, whereClause)

	var processedParams = processParamsForPostgreSQL(params...)
	builder.params = append(builder.params, processedParams...)

	return builder
}

func (builder *SQLTransactionalQueryBuilder[T]) Build() *SQLTransactionalQueryExecutor[T] {
	var processedSQL = processSQL(builder.querySQL, builder.whereClauses)
	builder.querySQL = processedSQL
	return &SQLTransactionalQueryExecutor[T]{
		queryBuilder: builder,
	}
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
//	executor := rdbms.BuildQuery[domain.Asset](adapter.GetDBX(), assetsSQL).Build()
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
	glog.Infof("Executing query %s \n with params %s", executor.query.SQL(), executor.query.Params())
	return executor.query.All(target)
}

func (executor *QueryExecutor[T]) GetInto(target *T) error {
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Executing query %s \n with params %s", executor.query.SQL(), executor.query.Params())
	return executor.query.One(target)
}

func (executor *QueryExecutor[T]) GetRows() (*dbx.Rows, error) {
	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Executing query %s \n with params %s", executor.query.SQL(), executor.query.Params())
	return executor.query.Rows()
}

// FindWithRowScanner executes the query and maps each result row with the provided scanner.
//
// Example:
//
//	queryExecutor := rdbms.BuildQuery[domain.Asset](adapter.GetDBX(), "SELECT id, ticker FROM asset").Build()
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
//	queryExecutor := rdbms.BuildQuery[domain.Asset](adapter.GetDBX(), "SELECT id, ticker FROM asset WHERE id = {:id}").
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

// ================================================
// ROW SCANNER
// ================================================

type RowScanner[T any] func(*sql.Rows) (T, error)

type SingleRowScanner[T any] func(*sql.Row) (T, error)

func ReturningIntIdRowScanner(rows *sql.Rows) (int64, error) {
	var id int64
	scanErr := rows.Scan(&id)
	if scanErr != nil {
		return 0, scanErr
	}
	return id, nil
}

func ReturningIntIdSingleRowScanner(row *sql.Row) (int64, error) {
	var id int64
	scanErr := row.Scan(&id)
	if scanErr != nil {
		return 0, scanErr
	}
	return id, nil
}

// ================================================
// QUERY EXECUTOR - sql.Tx
// ================================================

type SQLTransactionalQueryExecutor[T any] struct {
	queryBuilder *SQLTransactionalQueryBuilder[T]
}

func (executor *SQLTransactionalQueryExecutor[T]) Find(rowScanner RowScanner[T]) ([]T, error) {

	var builder = executor.queryBuilder

	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Executing transactional query %s \n with params %s", builder.querySQL, builder.params)
	rows, err := builder.transaction.Query(builder.querySQL, builder.params...)
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

		rowValue, scanErr := rowScanner(rows)
		if scanErr != nil {
			glog.Errorf("Error scanning row: %v", index)
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

func (executor *SQLTransactionalQueryExecutor[T]) Get(rowScanner SingleRowScanner[T]) (T, error) {

	var builder = executor.queryBuilder

	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Executing transactional query %s \n with params %s", builder.querySQL, builder.params)
	var row = builder.transaction.QueryRow(builder.querySQL, builder.params...)
	return rowScanner(row)
}
