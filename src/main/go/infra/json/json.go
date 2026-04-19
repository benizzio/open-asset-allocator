package json

import (
	"reflect"
	"strings"

	"github.com/benizzio/open-asset-allocator/langext"
)

// TODO verify if this functionality already exists in some library

// GetJSONFieldName converts a validator namespace into the equivalent JSON field path.
// It falls back to the provided fieldName when the namespace cannot be fully resolved.
//
// Co-authored by: OpenCode and GitHub Copilot
func GetJSONFieldName(namespace string, fieldName string, structType reflect.Type) string {
	var jsonFieldPath = buildJSONFieldPath(namespace, structType)
	if jsonFieldPath == "" {
		return fieldName
	}

	return jsonFieldPath
}

// buildJSONFieldPath resolves each namespace segment to its JSON field name while preserving
// collection indexes so validation messages match the external API contract.
//
// Authored by: OpenCode
func buildJSONFieldPath(namespace string, structType reflect.Type) string {
	var namespaceParts = parseNamespace(namespace)
	if len(namespaceParts) == 0 {
		return ""
	}

	var currentType = langext.UnwrapType(structType)
	if namespaceParts[0] == currentType.Name() {
		namespaceParts = namespaceParts[1:]
	}
	if len(namespaceParts) == 0 {
		return ""
	}

	var jsonParts = make([]string, 0, len(namespaceParts))
	for _, namespacePart := range namespaceParts {
		var jsonPart string
		var nextType reflect.Type
		if strings.Contains(namespacePart, "[") {
			jsonPart, nextType = mapIndexedNamespacePart(namespacePart, currentType)
		} else {
			jsonPart, nextType = mapNamespacePart(namespacePart, currentType)
		}
		if jsonPart == "" {
			return ""
		}

		jsonParts = append(jsonParts, jsonPart)
		currentType = nextType
	}

	return strings.Join(jsonParts, ".")
}

// mapNamespacePart resolves a single non-indexed validator namespace segment to its JSON name.
//
// Authored by: OpenCode
func mapNamespacePart(namespacePart string, currentType reflect.Type) (string, reflect.Type) {
	var structField, found = langext.FindStructFieldByNameOrJSONName(currentType, namespacePart)
	if !found {
		return "", nil
	}

	return langext.ExtractJSONFieldName(structField), langext.UnwrapType(structField.Type)
}

// mapIndexedNamespacePart resolves an indexed validator namespace segment like Allocations[0] to
// the corresponding JSON field path segment while preserving the original index.
//
// Authored by: OpenCode
func mapIndexedNamespacePart(namespacePart string, currentType reflect.Type) (string, reflect.Type) {
	var bracketIndex = strings.Index(namespacePart, "[")
	if bracketIndex == -1 {
		return "", nil
	}

	var fieldName = namespacePart[:bracketIndex]
	var structField, found = langext.FindStructFieldByNameOrJSONName(currentType, fieldName)
	if !found {
		return "", nil
	}

	var collectionType = langext.UnwrapType(structField.Type)
	if collectionType.Kind() != reflect.Array &&
		collectionType.Kind() != reflect.Slice &&
		collectionType.Kind() != reflect.Map {
		return "", nil
	}

	var jsonFieldName = langext.ExtractJSONFieldName(structField)
	var pathSuffix = namespacePart[bracketIndex:]
	return jsonFieldName + pathSuffix, langext.UnwrapType(collectionType.Elem())
}

// parseNamespace splits a namespace string into path segments.
//
// Co-authored by: OpenCode and GitHub Copilot
func parseNamespace(namespace string) []string {
	namespaceParts := strings.Split(namespace, ".")
	if len(namespaceParts) == 0 {
		return nil
	}

	return namespaceParts
}
