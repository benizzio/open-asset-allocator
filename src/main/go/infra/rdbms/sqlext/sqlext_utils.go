package sqlext

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/benizzio/open-asset-allocator/langext"
)

func ScanJsonColumn[T any](value interface{}, target *T) error {

	if value == nil {
		var zeroValue T
		*target = zeroValue
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scanned value is incompatible (not a []byte): %#v", value)
	}

	return json.Unmarshal(bytes, target)
}

// ValueJsonColumn serializes JSON-backed values as textual JSON so database/sql
// and pq.CopyIn use the same representation without coercing arbitrary binary
// payloads.
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
func ValueJsonColumn(value any) (driver.Value, error) {
	if value == nil {
		return nil, nil
	}

	if langext.IsNilPointer(value) {
		return nil, nil
	}

	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal %#v: %v", value, err)
	}

	return string(bytes), nil
}
