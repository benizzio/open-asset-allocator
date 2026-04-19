package rdbms

import (
	"database/sql"

	"github.com/golang/glog"
)

// ================================================
// QUERY BUILDER - sql.Tx
// ================================================

type SQLTransactionalQueryBuilder[T any] struct {
	transaction  *sql.Tx
	querySQL     string
	whereClauses []string
	params       []any
}

func (builder *SQLTransactionalQueryBuilder[T]) AddParams(
	params ...any,
) *SQLTransactionalQueryBuilder[T] {
	var processedParams = processParamsForPostgreSQL(params...)
	builder.params = append(builder.params, processedParams...)
	return builder
}

func (builder *SQLTransactionalQueryBuilder[T]) AddWhereClause(
	whereClause string,
) *SQLTransactionalQueryBuilder[T] {
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
// QUERY EXECUTOR - sql.Tx
// ================================================

type SQLTransactionalQueryExecutor[T any] struct {
	queryBuilder *SQLTransactionalQueryBuilder[T]
}

func (executor *SQLTransactionalQueryExecutor[T]) Find(rowScanner RowScanner[T]) ([]T, error) {

	var builder = executor.queryBuilder

	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Executing transactional query %s \n with params %v", builder.querySQL, builder.params)
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

func (executor *SQLTransactionalQueryExecutor[T]) Get(rowScanner SingleRowScanner[T]) (T, error) {

	var builder = executor.queryBuilder

	// TODO verification for debug logging, this should be logged only in debug mode
	glog.Infof("Executing transactional query %s \n with params %v", builder.querySQL, builder.params)
	var row = builder.transaction.QueryRow(builder.querySQL, builder.params...)
	return rowScanner(row)
}
