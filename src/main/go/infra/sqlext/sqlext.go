package sqlext

import (
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/lib/pq"
)

type NullStringSlice []sql.NullString

func (nullStringSlice *NullStringSlice) Scan(src interface{}) error {
	if src == nil {
		*nullStringSlice = []sql.NullString{}
		return nil
	}
	return pq.Array(nullStringSlice).Scan(src)
}

func (nullStringSlice NullStringSlice) Value() (driver.Value, error) {
	return pq.Array(nullStringSlice).Value()
}

func (nullStringSlice *NullStringSlice) ToStringSlice() []*string {
	var result = make([]*string, 0)
	for _, item := range *nullStringSlice {
		var itemReference *string
		if item.Valid {
			itemReference = &item.String
		}
		result = append(result, itemReference)
	}
	return result
}

// BuildNullStringSlice constructs a NullStringSlice from a slice of string pointers,
// preserving nil entries as sql.NullString with Valid=false. Useful to encode
// PostgreSQL text[] parameters via pq.Array while retaining NULL elements.
//
// Parameters:
//   - values: slice of string pointers where nil indicates SQL NULL
//
// Returns:
//   - NullStringSlice: slice ready to be wrapped by pq.Array for database/sql usage
//
// Authored by: GitHub Copilot
func BuildNullStringSlice(values []*string) NullStringSlice {
	var arr = make(NullStringSlice, len(values))
	for i, ptr := range values {
		if ptr == nil {
			arr[i] = sql.NullString{String: "", Valid: false}
		} else {
			arr[i] = sql.NullString{String: *ptr, Valid: true}
		}
	}
	return arr
}

type NullTime sql.NullTime

func (nullTime *NullTime) ToTimeReference() *time.Time {
	if nullTime.Valid {
		return &nullTime.Time
	}
	return nil
}
