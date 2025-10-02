package util

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

func StringToNullString(str string) sql.NullString {
	return StringPointerToNullString(&str)
}

func StringPointerToNullString(str *string) sql.NullString {
	if str == nil {
		return sql.NullString{Valid: false}
	} else {
		return sql.NullString{String: *str, Valid: true}
	}
}

// ValueToString normalizes driver.Value to a comparable string.
// It handles string and []byte, and falls back to fmt.Sprintf otherwise.
//
// Co-authored by: GitHub Copilot
func ValueToString(value driver.Value) string {
	switch valueType := value.(type) {
	case string:
		return valueType
	case []byte:
		return string(valueType)
	default:
		return fmt.Sprintf("%v", value)
	}
}
