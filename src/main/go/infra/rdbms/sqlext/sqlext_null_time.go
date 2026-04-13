package sqlext

import (
	"database/sql"
	"time"
)

type NullTime sql.NullTime

func (nullTime *NullTime) ToTimeReference() *time.Time {
	if nullTime == nil || !nullTime.Valid {
		return nil
	}

	return new(nullTime.Time)
}
