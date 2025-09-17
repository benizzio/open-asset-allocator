package langext

import (
	"encoding/json"
	"errors"
	"strconv"
)

// ParseableInt64 provides JSON marshaling/unmarshaling for int64 values that can be either
// numeric or string in JSON.
//
// Authored by: GitHub Copilot
type ParseableInt64 int64

// UnmarshalJSON implements json.Unmarshaler interface for ParseableInt64.
// It can parse both numeric and string representations of int64 values.
//
// Authored by: GitHub Copilot
func (parseableInt64 *ParseableInt64) UnmarshalJSON(data []byte) error {

	var value int64
	if err := json.Unmarshal(data, &value); err == nil {
		*parseableInt64 = ParseableInt64(value)
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

	var parsedValue int64
	var err error
	if IsZeroValue(stringValue) {
		parsedValue = 0
	} else {
		parsedValue, err = strconv.ParseInt(stringValue, 10, 64)
	}
	if err != nil {
		return err
	}

	*parseableInt64 = ParseableInt64(parsedValue)
	return nil
}