package langext

import (
	"encoding/json"
	"errors"
	"strconv"
)

type ParseableInt int

func (parseableInt *ParseableInt) UnmarshalJSON(data []byte) error {

	var value int
	if err := json.Unmarshal(data, &value); err == nil {
		*parseableInt = ParseableInt(value)
		return nil
	} else {
		// Check if the error is specifically a type error
		var unmarshalTypeError *json.UnmarshalTypeError
		if !errors.As(err, &unmarshalTypeError) {
			// If it's not a type error, return the original error
			return err
		}
	}

	// Fallback to convert if it's a string
	var stringValue string
	if err := json.Unmarshal(data, &stringValue); err != nil {
		return err
	}

	parsedValue, err := strconv.Atoi(stringValue)
	if err != nil {
		return err
	}
	*parseableInt = ParseableInt(parsedValue)
	return nil
}
