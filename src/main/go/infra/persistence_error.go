package infra

import "fmt"

// ForeignKeyConstraintViolationError identifies a rejected reference at the persistence boundary.
// Domain services may translate a known constraint into a domain validation error without importing a database driver.
// Authored by: GPT-6.1 Sol
type ForeignKeyConstraintViolationError struct {
	ConstraintName string
	Cause          error
}

// Error describes the violated persistence constraint.
// For example, fmt.Sprint(err) includes the constraint name for diagnostics.
// Authored by: GPT-6.1 Sol
func (constraintError *ForeignKeyConstraintViolationError) Error() string {
	return "Foreign key constraint violation: " + constraintError.ConstraintName
}

// Unwrap exposes the database cause for standard error inspection, such as errors.As(err, &driverError).
// Authored by: GPT-6.1 Sol
func (constraintError *ForeignKeyConstraintViolationError) Unwrap() error {
	return constraintError.Cause
}

// UnexpectedRowsAffectedError reports a persistence write that did not affect its expected number of rows.
// It is a persistence failure, not a domain validation decision.
// Authored by: GPT-6.1 Sol
type UnexpectedRowsAffectedError struct {
	Operation string
	Expected  int64
	Actual    int64
}

// Error describes the write and its actual row count, for example when logging an update failure.
// Authored by: GPT-6.1 Sol
func (persistenceError *UnexpectedRowsAffectedError) Error() string {
	return fmt.Sprintf("%s: expected %d affected rows, got %d",
		persistenceError.Operation, persistenceError.Expected, persistenceError.Actual)
}
