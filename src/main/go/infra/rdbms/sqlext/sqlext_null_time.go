package sqlext

import (
	"database/sql"
	"time"
)

type NullTime sql.NullTime

func (nullTime *NullTime) ToTimeReference() *time.Time {
	if nullTime.Valid {
		return &nullTime.Time
	}
	return nil
}
