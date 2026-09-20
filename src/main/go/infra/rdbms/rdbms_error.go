// rdbms_error.go contains database error classification helpers shared by RDBMS repositories.
//
// Authored by: OpenCode
package rdbms

import (
	"errors"

	"github.com/lib/pq"
)

// IsUniqueConstraintViolation reports whether an RDBMS error represents a unique-constraint
// violation for the supplied constraint name.
//
// Example:
//
//	if IsUniqueConstraintViolation(err, "asset_ticker_uk") {
//		// Map the database error to the repository's application error.
//	}
//
// Authored by: OpenCode
func IsUniqueConstraintViolation(err error, constraintName string) bool {
	var postgresError *pq.Error
	return errors.As(err, &postgresError) &&
		postgresError.Code == "23505" &&
		postgresError.Constraint == constraintName
}
