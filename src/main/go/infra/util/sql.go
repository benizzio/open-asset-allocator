package util

import "database/sql"

func ToNullString(str string) sql.NullString {
	return ToNullStringFromPointer(&str)
}

func ToNullStringFromPointer(str *string) sql.NullString {
	if str == nil {
		return sql.NullString{Valid: false}
	} else {
		return sql.NullString{String: *str, Valid: true}
	}
}
