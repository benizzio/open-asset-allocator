package util

import (
	"reflect"
	"strings"
)

// getJSONFieldName returns the field name from the JSON tag of the struct field.
//
// Authored by: GitHub Copilot
func getJSONFieldName(namespace string, fieldName string, structType reflect.Type) string {
	// Get the struct field
	field := findFieldByNamespace(namespace, structType)
	if field == nil {
		// Fall back to the provided field name
		return fieldName
	}

	// Get the json tag and parse it
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		// No json tag found, return original field name
		return fieldName
	}

	// The json tag might have options like ",omitempty", so we need to split it
	jsonName := strings.Split(jsonTag, ",")[0]
	if jsonName == "-" || jsonName == "" {
		// Field is ignored in JSON or empty tag, return original field name
		return fieldName
	}

	return jsonName
}

// findFieldByNamespace navigates through a struct to find a field using the namespace.
//
// Authored by: GitHub Copilot
func findFieldByNamespace(namespace string, structType reflect.Type) *reflect.StructField {
	namespaceParts := parseNamespace(namespace)
	if len(namespaceParts) == 0 {
		return nil
	}

	return navigateToField(namespaceParts, structType)
}

// parseNamespace splits a namespace string into parts, skipping the struct name.
//
// Authored by: GitHub Copilot
func parseNamespace(namespace string) []string {
	namespaceParts := strings.Split(namespace, ".")
	if len(namespaceParts) < 2 {
		return nil // Invalid namespace format
	}

	// Skip the first part which is the struct name
	return namespaceParts[1:]
}

// navigateToField follows the path through a struct to find a specific field.
//
// Authored by: GitHub Copilot
func navigateToField(fieldPath []string, startType reflect.Type) *reflect.StructField {
	// Return early if path is empty
	if len(fieldPath) == 0 {
		return nil
	}

	// Process recursively starting from the first field
	return findNestedField(fieldPath, 0, startType)
}

// findNestedField recursively processes a field path to find the target field.
//
// Authored by: GitHub Copilot
func findNestedField(fieldPath []string, pathIndex int, currentType reflect.Type) *reflect.StructField {
	// We've reached the end of the path
	if pathIndex >= len(fieldPath) {
		return nil
	}

	currentFieldName := fieldPath[pathIndex]
	isLastField := pathIndex == len(fieldPath)-1

	// Process the current field based on its type
	if strings.Contains(currentFieldName, "[") {
		return processArrayField(currentFieldName, fieldPath, pathIndex, isLastField, currentType)
	}

	return processRegularField(currentFieldName, fieldPath, pathIndex, isLastField, currentType)
}

// processArrayField handles array/slice fields in the path.
//
// Authored by: GitHub Copilot
func processArrayField(
	fieldPathComponent string,
	fieldPath []string,
	pathIndex int,
	isLastField bool,
	currentType reflect.Type,
) *reflect.StructField {
	structField, elementType := handleArrayField(fieldPathComponent, currentType)
	if structField == nil {
		return nil
	}

	// If this is the last part, return the field
	if isLastField {
		return structField
	}

	// Otherwise, continue navigating with the element type
	return findNestedField(fieldPath, pathIndex+1, elementType)
}

// processRegularField handles regular struct fields in the path.
//
// Authored by: GitHub Copilot
func processRegularField(
	fieldPathComponent string,
	fieldPath []string,
	pathIndex int,
	isLastField bool,
	currentType reflect.Type,
) *reflect.StructField {
	structField, found := currentType.FieldByName(fieldPathComponent)
	if !found {
		return nil
	}

	// If this is the last part, return the field
	if isLastField {
		return &structField
	}

	// Otherwise, continue navigating with the field's type
	nextType := unwrapType(structField.Type)
	return findNestedField(fieldPath, pathIndex+1, nextType)
}

// handleArrayField processes array/slice field references and returns the field and element type.
// Returns nil values when the field can't be found or is not an array/slice.
//
// Authored by: GitHub Copilot
func handleArrayField(fieldPathComponent string, currentType reflect.Type) (*reflect.StructField, reflect.Type) {
	// Extract the field name before the array index
	bracketIndex := strings.Index(fieldPathComponent, "[")
	if bracketIndex == -1 {
		return nil, nil
	}

	arrayFieldName := fieldPathComponent[:bracketIndex]

	structField, found := currentType.FieldByName(arrayFieldName)
	if !found {
		return nil, nil
	}

	// Get the element type for arrays/slices/maps
	fieldType := unwrapType(structField.Type)
	var elementType reflect.Type

	switch fieldType.Kind() {
	case reflect.Array, reflect.Slice:
		elementType = fieldType.Elem()
	case reflect.Map:
		elementType = fieldType.Elem()
	default:
		// Not an indexable type
		return nil, nil
	}

	return &structField, elementType
}
