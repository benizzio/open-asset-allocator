package rdbms

import (
	"database/sql"
	"strings"

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
