package langext

import (
	"encoding/json"
	"reflect"
	"strings"
)

// DeepCompleteStruct fills in zero values in the target struct with values from the source struct.
func DeepCompleteStruct[T interface{}](target *T, source *T) {

	if target == nil || source == nil {
		return
	}

	var structType = reflect.TypeOf(*target)
	var targetStructValue = reflect.ValueOf(target).Elem()
	var sourceStructValue = reflect.ValueOf(source).Elem()

	if structType.Kind() == reflect.Struct {
		deepCompleteReflective(structType, targetStructValue, sourceStructValue)
	}
}

func deepCompleteReflective(structType reflect.Type, targetStructValue reflect.Value, sourceStructValue reflect.Value) {

	if !targetStructValue.IsZero() {
		return
	}

	for i := 0; i < structType.NumField(); i++ {

		if structType.Field(i).Type.Kind() == reflect.Struct {
			deepCompleteReflective(
				structType.Field(i).Type,
				targetStructValue.Field(i).Elem(),
				sourceStructValue.Field(i).Elem(),
			)
		}

		var sourceStructFieldValue = sourceStructValue.Field(i)
		var targetStructFieldValue = targetStructValue.Field(i)

		if targetStructFieldValue.IsZero() {
			targetStructFieldValue.Set(sourceStructFieldValue)
		}
	}
}

// StructString converts a struct to its JSON string representation.
func StructString[T interface{}](source *T) string {
	out, err := json.Marshal(source)
	if err != nil {
		panic(err)
	}
	return string(out)
}

// GetStructName extracts the struct name from a struct or pointer to struct.
// It handles both struct values and pointers to structs.
//
// Parameters:
//   - targetStruct: The struct or pointer to struct to get the name from
//
// Returns:
//   - The name of the struct type
//
// Authored by: GitHub Copilot
func GetStructName(targetStruct interface{}) string {
	structType := reflect.TypeOf(targetStruct)

	// Handle pointer types
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	return structType.Name()
}

// GetStructNamespaceDescription extracts field name and builds the full namespace
// from a struct and field namespace string.
//
// Parameters:
//   - targetStruct: The struct or pointer to struct
//   - fieldNamespace: The field namespace (can be a simple field name or a dot-separated path)
//
// Returns:
//   - namespace: The full namespace including the struct name
//   - fieldName: The simple field name (last part of the namespace)
//
// Authored by: GitHub Copilot
func GetStructNamespaceDescription(targetStruct interface{}, fieldNamespace string) (namespace, fieldName string) {
	var structName = GetStructName(targetStruct)
	var parts = strings.Split(fieldNamespace, ".")

	// The field name is always the last part of the namespace
	fieldName = parts[len(parts)-1]

	// If namespace doesn't already start with struct name, prepend it
	if !strings.HasPrefix(fieldNamespace, structName+".") {
		namespace = structName + "." + fieldNamespace
	} else {
		namespace = fieldNamespace
	}

	return namespace, fieldName
}
