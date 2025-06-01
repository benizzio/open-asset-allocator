package util

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// arrayIndexPattern is the regex pattern used to identify array index notation
//
// Authored by: GitHub Copilot
var arrayIndexPattern = regexp.MustCompile(`^(.+)\[(\d+)]$`)

// AssertJSONEqualIgnoringFields compares two JSON strings for equality after removing specified fields.
//
// fieldsToIgnore supports various path notations to identify fields to ignore:
//
// 1. Simple field names for top-level fields:
//   - "id" -> removes the top-level "id" field
//
// 2. Dot notation for nested fields:
//   - "user.name" -> removes the "name" field nested inside the "user" object
//   - "meta.data.created" -> removes the "created" field nested in "data" which is in "meta"
//
// 3. Array traversal (automatic):
//   - "items.name" -> removes the "name" field from ALL items in the array
//   - "data.users.addresses.street" -> removes the "street" field from all addresses
//     of all users
//
// 4. Array notation with specific indexes (optional):
//   - "items[0]" -> removes the first element of the "items" array
//   - "users[2].name" -> removes the "name" field from only the third user in the array
//
// Examples:
//
//	assertJSONEqualIgnoringFields(t, expected, actual, "id")
//	assertJSONEqualIgnoringFields(t, expected, actual, "user.createdAt", "meta.version")
//	assertJSONEqualIgnoringFields(t, expected, actual, "allocations.id", "history.timestamp")
//
// Authored by: GitHub Copilot
func AssertJSONEqualIgnoringFields(t *testing.T, expectedJSON, actualJSON string, fieldsToIgnore ...string) {
	// Parse expected JSON to a map
	var expectedMap map[string]interface{}
	err := json.Unmarshal([]byte(expectedJSON), &expectedMap)
	assert.NoError(t, err, "Failed to parse expected JSON")

	// Parse actual JSON to a map
	var actualMap map[string]interface{}
	err = json.Unmarshal([]byte(actualJSON), &actualMap)
	assert.NoError(t, err, "Failed to parse actual JSON")

	// Remove specified fields from both maps
	for _, field := range fieldsToIgnore {
		// For top-level fields, handle as before
		if !strings.Contains(field, ".") {
			delete(expectedMap, field)
			delete(actualMap, field)
		} else {
			// For nested fields, use the helper function
			removeNestedField(expectedMap, field)
			removeNestedField(actualMap, field)
		}
	}

	// Convert maps back to JSON
	expectedWithoutFields, err := json.Marshal(expectedMap)
	assert.NoError(t, err, "Failed to marshal modified expected JSON")

	actualWithoutFields, err := json.Marshal(actualMap)
	assert.NoError(t, err, "Failed to marshal modified actual JSON")

	// Compare JSONs after removing specified fields
	assert.JSONEq(t, string(expectedWithoutFields), string(actualWithoutFields))
}

// removeNestedField removes a nested field specified by a dot-notation path.
// The path can include array notation with numeric indices in square brackets,
// or simple dot notation which will automatically apply to all elements when an array is encountered.
// For example:
// - "users[0].address.street" would target the street field in the first user's address
// - "users.address.street" would target the street field in all users' addresses
//
// Authored by: GitHub Copilot
func removeNestedField(data map[string]interface{}, fieldPath string) {
	parts := strings.Split(fieldPath, ".")
	removeNestedFieldRecursive(data, parts, 0)
}

// removeNestedFieldRecursive recursively traverses the data structure to remove a field.
// It handles two types of path segments:
// 1. Regular field names (e.g., "user", "address", "items")
// 2. Array access with specific index (e.g., "items[0]", "users[2]")
//
// When an array is encountered during traversal with a regular field name,
// it automatically applies the operation to all elements in the array.
//
// Authored by: GitHub Copilot
func removeNestedFieldRecursive(data interface{}, parts []string, depth int) {
	// Base case: we've reached the end of our path
	if depth >= len(parts)-1 {
		return
	}

	currentPart := parts[depth]
	nextPart := parts[depth+1]
	isLastLevel := depth == len(parts)-2

	// Check if we're processing an array index notation or a regular field
	if isArrayIndexNotation(currentPart) {
		handleArrayIndexPath(data, currentPart, parts, depth, isLastLevel)
		return
	}

	// Handle regular field case
	handleRegularFieldPath(data, currentPart, nextPart, parts, depth, isLastLevel)
}

// isArrayIndexNotation checks if a path segment uses array index notation like "field[0]"
//
// Authored by: GitHub Copilot
func isArrayIndexNotation(pathSegment string) bool {
	return len(arrayIndexPattern.FindStringSubmatch(pathSegment)) > 0
}

// extractArrayIndexInfo extracts the field name and index from an array index path segment
//
// Authored by: GitHub Copilot
func extractArrayIndexInfo(pathSegment string) (string, int) {
	matches := arrayIndexPattern.FindStringSubmatch(pathSegment)
	if len(matches) == 0 {
		return "", -1
	}
	fieldName := matches[1]
	index, _ := strconv.Atoi(matches[2])
	return fieldName, index
}

// handleArrayIndexPath processes a path segment that contains array index notation
//
// Authored by: GitHub Copilot
func handleArrayIndexPath(data interface{}, currentPart string, parts []string, depth int, isLastLevel bool) {
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return
	}

	fieldName, index := extractArrayIndexInfo(currentPart)
	arr, ok := dataMap[fieldName].([]interface{})
	if !ok || index >= len(arr) {
		return // Array doesn't exist or index out of bounds
	}

	if isLastLevel {
		handleFinalRemoval(arr[index], parts[depth+1])
	} else {
		removeNestedFieldRecursive(arr[index], parts, depth+1)
	}
}

// handleRegularFieldPath processes a regular field path segment
//
// Authored by: GitHub Copilot
func handleRegularFieldPath(
	data interface{},
	currentPart, nextPart string,
	parts []string,
	depth int,
	isLastLevel bool,
) {
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return
	}

	currentValue, exists := dataMap[currentPart]
	if !exists {
		return // Field doesn't exist
	}

	// Check if we need to process all elements of an array
	if arr, isArray := currentValue.([]interface{}); isArray {
		processArrayElements(arr, parts, depth, isLastLevel)
		return
	}

	// Process a single value
	if isLastLevel {
		handleFinalRemoval(currentValue, nextPart)
	} else {
		removeNestedFieldRecursive(currentValue, parts, depth+1)
	}
}

// processArrayElements applies the operation to all elements in an array
//
// Authored by: GitHub Copilot
func processArrayElements(arr []interface{}, parts []string, depth int, isLastLevel bool) {
	nextPart := parts[depth+1]

	for i := 0; i < len(arr); i++ {
		if isLastLevel {
			handleFinalRemoval(arr[i], nextPart)
		} else {
			removeNestedFieldRecursive(arr[i], parts, depth+1)
		}
	}
}

// handleFinalRemoval removes the specified field from the final object in the path.
// It supports specific array indices in the final segment and auto-detects arrays
// to apply removal to all elements.
//
// Authored by: GitHub Copilot
func handleFinalRemoval(data interface{}, finalPart string) {
	// Check if final part contains array notation
	arraySpecificMatches := arrayIndexPattern.FindStringSubmatch(finalPart)

	// Handle specific array index in final part
	if len(arraySpecificMatches) > 0 {
		fieldName := arraySpecificMatches[1]
		index, _ := strconv.Atoi(arraySpecificMatches[2])

		if valueMap, ok := data.(map[string]interface{}); ok {
			if arr, ok := valueMap[fieldName].([]interface{}); ok && index < len(arr) {
				// Create a new array without the specified element
				newArr := append(arr[:index], arr[index+1:]...)
				valueMap[fieldName] = newArr
			}
		}
		return
	}

	// Handle regular field deletion
	if valueMap, ok := data.(map[string]interface{}); ok {
		// Auto-detect if the field is an array
		if _, isArray := valueMap[finalPart].([]interface{}); isArray {
			// If the field is an array, empty it
			valueMap[finalPart] = []interface{}{}
			return
		}

		// Otherwise delete the field normally
		delete(valueMap, finalPart)
	}
}
