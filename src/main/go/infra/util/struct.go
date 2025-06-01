package util

import (
	"encoding/json"
	"reflect"
)

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
func GetStructName(targetStruct interface{}) string {
	structType := reflect.TypeOf(targetStruct)

	// Handle pointer types
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	return structType.Name()
}
