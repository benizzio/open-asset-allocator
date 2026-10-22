// rdbms_error.go contains database error classification helpers shared by RDBMS repositories.
//
// Co-authored by: GPT-6.1 Sol and OpenCode
package rdbms

import (
	"errors"

	"github.com/lib/pq"

	"github.com/benizzio/open-asset-allocator/infra"
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

// ClassifyForeignKeyViolation returns persistence metadata for any foreign-key violation,
// including wrapped driver errors, or nil for other errors. It preserves the constraint name
// and original cause so callers can interpret known relationships without importing the driver.
//
// Example:
//
//	if violation := ClassifyForeignKeyViolation(err); violation != nil {
//		return violation
//	}
//
// Authored by: GPT-6.1 Sol
func ClassifyForeignKeyViolation(err error) *infra.ForeignKeyConstraintViolationError {
	var postgresError *pq.Error
	if !errors.As(err, &postgresError) || postgresError.Code != "23503" {
		return nil
	}
	return &infra.ForeignKeyConstraintViolationError{ConstraintName: postgresError.Constraint, Cause: err}
}
