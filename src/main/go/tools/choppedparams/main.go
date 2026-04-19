// Package main validates that multiline Go calls and function declarations keep one parameter per line.
//
// Authored by: GitHub Copilot
package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

// violation describes a single chopped-parameter formatting failure.
//
// Authored by: GitHub Copilot
type violation struct {
	file    string
	line    int
	column  int
	message string
}

// main runs the chopped-parameter validator for the provided files and directories.
//
// Authored by: GitHub Copilot
func main() {
	err := run(os.Args[1:])
	if err == nil {
		return
	}

	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

// run executes the chopped-parameter validation workflow.
//
// Authored by: GitHub Copilot
func run(paths []string) error {
	files, err := collectGoFiles(paths)
	if err != nil {
		return err
	}

	violations, err := collectViolations(files)
	if err != nil {
		return err
	}

	if len(violations) == 0 {
		return nil
	}

	for _, item := range violations {
		_, _ = fmt.Fprintf(
			os.Stderr,
			"%s:%d:%d: %s\n",
			item.file,
			item.line,
			item.column,
			item.message,
		)
	}

	return fmt.Errorf("found %d multiline chopped-parameter violation(s)", len(violations))
}

// collectGoFiles expands the provided file and directory arguments into Go source files.
//
// Authored by: GitHub Copilot
func collectGoFiles(paths []string) ([]string, error) {
	if len(paths) == 0 {
		return nil, errors.New("no paths provided")
	}

	var fileSet = map[string]struct{}{}

	for _, inputPath := range paths {
		err := filepath.WalkDir(inputPath, func(path string, entry fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}

			if entry.IsDir() {
				if shouldSkipDirectory(path, entry.Name()) {
					return filepath.SkipDir
				}

				return nil
			}

			if filepath.Ext(path) != ".go" {
				return nil
			}

			fileSet[filepath.Clean(path)] = struct{}{}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	var files = make([]string, 0, len(fileSet))
	for file := range fileSet {
		files = append(files, file)
	}

	slices.Sort(files)
	return files, nil
}

// shouldSkipDirectory filters directories that should not be linted recursively.
//
// Authored by: GitHub Copilot
func shouldSkipDirectory(path string, name string) bool {
	if path == "." {
		return false
	}

	return name == ".git" || name == "target" || name == "vendor"
}

// collectViolations parses the provided Go files and gathers formatting violations.
//
// Authored by: GitHub Copilot
func collectViolations(files []string) ([]violation, error) {
	var fileSet = token.NewFileSet()
	var violations []violation

	for _, filePath := range files {
		fileViolations, err := validateGoFile(fileSet, filePath)
		if err != nil {
			return nil, err
		}

		violations = append(violations, fileViolations...)
	}

	slices.SortFunc(violations, compareViolations)
	return violations, nil
}

// validateGoFile scans a single Go source file for chopped-parameter violations.
//
// Authored by: GitHub Copilot
func validateGoFile(fileSet *token.FileSet, filePath string) ([]violation, error) {
	fileNode, err := parser.ParseFile(fileSet, filePath, nil, 0)
	if err != nil {
		return nil, err
	}

	var violations []violation

	ast.Inspect(fileNode, func(node ast.Node) bool {
		switch typedNode := node.(type) {
		case *ast.CallExpr:
			violations = append(violations, validateCallArguments(fileSet, filePath, typedNode)...)
		case *ast.FuncType:
			violations = append(violations, validateFunctionParameters(fileSet, filePath, typedNode.Params)...)
		}

		return true
	})

	return violations, nil
}

// validateCallArguments checks that multiline call arguments stay one per line.
//
// Authored by: GitHub Copilot
func validateCallArguments(fileSet *token.FileSet, filePath string, callExpression *ast.CallExpr) []violation {
	if len(callExpression.Args) < 2 || !spansMultipleLines(fileSet, callExpression.Lparen, callExpression.Rparen) {
		return nil
	}

	var violations []violation
	var openingLine = fileSet.Position(callExpression.Lparen).Line
	var firstArgumentLine = fileSet.Position(callExpression.Args[0].Pos()).Line

	if firstArgumentLine == openingLine {
		violations = append(
			violations,
			buildViolation(
				fileSet,
				filePath,
				callExpression.Args[0].Pos(),
				"multiline call arguments must start on their own lines",
			),
		)
	}

	var seenLines = map[int]struct{}{}
	for _, argument := range callExpression.Args {
		var argumentPosition = fileSet.Position(argument.Pos())

		if _, exists := seenLines[argumentPosition.Line]; exists {
			violations = append(
				violations,
				buildViolation(
					fileSet,
					filePath,
					argument.Pos(),
					"multiline call arguments must be chopped to one per line",
				),
			)
			break
		}

		seenLines[argumentPosition.Line] = struct{}{}
	}

	return violations
}

// validateFunctionParameters checks that multiline function parameters stay one per line.
//
// Authored by: GitHub Copilot
func validateFunctionParameters(fileSet *token.FileSet, filePath string, parameters *ast.FieldList) []violation {
	if parameters == nil || parameterCount(parameters) < 2 {
		return nil
	}

	if !spansMultipleLines(fileSet, parameters.Opening, parameters.Closing) {
		return nil
	}

	var violations []violation
	var openingLine = fileSet.Position(parameters.Opening).Line
	var firstParameterPosition = fieldPosition(parameters.List[0])
	var firstParameterLine = fileSet.Position(firstParameterPosition).Line

	if firstParameterLine == openingLine {
		violations = append(
			violations,
			buildViolation(
				fileSet,
				filePath,
				firstParameterPosition,
				"multiline function parameters must start on their own lines",
			),
		)
	}

	var seenLines = map[int]struct{}{}
	for _, field := range parameters.List {
		var position = fieldPosition(field)
		var line = fileSet.Position(position).Line

		if len(field.Names) > 1 {
			violations = append(
				violations,
				buildViolation(
					fileSet,
					filePath,
					field.Names[1].Pos(),
					"multiline function parameters must be chopped to one per line",
				),
			)
		}

		if _, exists := seenLines[line]; exists {
			violations = append(
				violations,
				buildViolation(
					fileSet,
					filePath,
					position,
					"multiline function parameters must be chopped to one per line",
				),
			)
			break
		}

		seenLines[line] = struct{}{}
	}

	return violations
}

// parameterCount counts the logical number of declared function parameters.
//
// Authored by: GitHub Copilot
func parameterCount(parameters *ast.FieldList) int {
	var total = 0

	for _, field := range parameters.List {
		if len(field.Names) == 0 {
			total++
			continue
		}

		total += len(field.Names)
	}

	return total
}

// fieldPosition returns the most specific position available for a field declaration.
//
// Authored by: GitHub Copilot
func fieldPosition(field *ast.Field) token.Pos {
	if len(field.Names) > 0 {
		return field.Names[0].Pos()
	}

	return field.Pos()
}

// spansMultipleLines reports whether a token range crosses more than one source line.
//
// Authored by: GitHub Copilot
func spansMultipleLines(fileSet *token.FileSet, opening token.Pos, closing token.Pos) bool {
	if !opening.IsValid() || !closing.IsValid() {
		return false
	}

	return fileSet.Position(opening).Line != fileSet.Position(closing).Line
}

// buildViolation creates a violation from a token position and message.
//
// Authored by: GitHub Copilot
func buildViolation(fileSet *token.FileSet, filePath string, position token.Pos, message string) violation {
	var sourcePosition = fileSet.Position(position)

	return violation{
		file:    filePath,
		line:    sourcePosition.Line,
		column:  sourcePosition.Column,
		message: message,
	}
}

// compareViolations keeps violation output stable across runs.
//
// Authored by: GitHub Copilot
func compareViolations(left violation, right violation) int {
	if left.file != right.file {
		return compareStrings(left.file, right.file)
	}

	if left.line != right.line {
		return left.line - right.line
	}

	if left.column != right.column {
		return left.column - right.column
	}

	return compareStrings(left.message, right.message)
}

// compareStrings compares two strings for sorting.
//
// Authored by: GitHub Copilot
func compareStrings(left string, right string) int {
	if left < right {
		return -1
	}

	if left > right {
		return 1
	}

	return 0
}
