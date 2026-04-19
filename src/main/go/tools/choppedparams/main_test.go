// Package main tests the chopped-parameter validator behavior.
//
// Authored by: GitHub Copilot
package main

import (
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestValidateGoFileAllowsSingleLineAndChoppedMultilineCode verifies accepted formatting cases.
//
// Authored by: GitHub Copilot
func TestValidateGoFileAllowsSingleLineAndChoppedMultilineCode(t *testing.T) {
	t.Parallel()

	var filePath = writeTestFile(
		t,
		"allowed.go",
		`package sample

func allowedSingleLine(a string, b string) {}

func allowedMultiline(
	a string,
	b string,
) {
	call(
		firstArg,
		secondArg,
	)
}
`,
	)

	var fileSet = token.NewFileSet()
	violations, err := validateGoFile(fileSet, filePath)
	if err != nil {
		t.Fatalf("validateGoFile returned an unexpected error: %v", err)
	}

	if len(violations) != 0 {
		t.Fatalf("expected no violations, got %d", len(violations))
	}
}

// TestValidateGoFileRejectsMultilineCallsWithoutChoppedArguments verifies multiline call enforcement.
//
// Authored by: GitHub Copilot
func TestValidateGoFileRejectsMultilineCallsWithoutChoppedArguments(t *testing.T) {
	t.Parallel()

	var filePath = writeTestFile(
		t,
		"call_violation.go",
		`package sample

func broken() {
	call(
		firstArg, secondArg,
	)
}
`,
	)

	var fileSet = token.NewFileSet()
	violations, err := validateGoFile(fileSet, filePath)
	if err != nil {
		t.Fatalf("validateGoFile returned an unexpected error: %v", err)
	}

	if len(violations) == 0 {
		t.Fatal("expected at least one violation")
	}

	if !containsMessage(violations, "multiline call arguments must be chopped to one per line") {
		t.Fatalf("expected a multiline call violation, got %#v", violations)
	}
}

// TestValidateGoFileRejectsMultilineFunctionsWithoutChoppedParameters verifies multiline declaration enforcement.
//
// Authored by: GitHub Copilot
func TestValidateGoFileRejectsMultilineFunctionsWithoutChoppedParameters(t *testing.T) {
	t.Parallel()

	var filePath = writeTestFile(
		t,
		"func_violation.go",
		`package sample

func broken(
	firstParam, secondParam string,
) {
}
`,
	)

	var fileSet = token.NewFileSet()
	violations, err := validateGoFile(fileSet, filePath)
	if err != nil {
		t.Fatalf("validateGoFile returned an unexpected error: %v", err)
	}

	if len(violations) == 0 {
		t.Fatal("expected at least one violation")
	}

	if !containsMessage(violations, "multiline function parameters must be chopped to one per line") {
		t.Fatalf("expected a multiline function parameter violation, got %#v", violations)
	}
}

// writeTestFile creates a temporary Go file for validation tests.
//
// Authored by: GitHub Copilot
func writeTestFile(t *testing.T, fileName string, content string) string {
	t.Helper()

	var directory = t.TempDir()
	var filePath = filepath.Join(directory, fileName)
	err := os.WriteFile(filePath, []byte(content), 0o600)
	if err != nil {
		t.Fatalf("writeTestFile failed: %v", err)
	}

	return filePath
}

// containsMessage checks whether the violation list contains the expected message fragment.
//
// Authored by: GitHub Copilot
func containsMessage(violations []violation, message string) bool {
	for _, violation := range violations {
		if strings.Contains(violation.message, message) {
			return true
		}
	}

	return false
}
