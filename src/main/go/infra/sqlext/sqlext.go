package sqlext

import (
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type NullStringSlice []sql.NullString

func (nullStringSlice *NullStringSlice) Scan(src interface{}) error {
	if src == nil {
		*nullStringSlice = []sql.NullString{}
		return nil
	}
	return pq.Array(nullStringSlice).Scan(src)
}

func (nullStringSlice *NullStringSlice) ToStringSlice() []string {
	var result = make([]string, 0)
	for _, item := range *nullStringSlice {
		if item.Valid {
			result = append(result, item.String)
		}
	}
	return result
}

type NullTime sql.NullTime

func (nullTime *NullTime) ToTimeReference() *time.Time {
	if nullTime.Valid {
		return &nullTime.Time
	}
	return nil
}
