package sqlext

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

func ScanJsonColumn[T any](value interface{}, target *T) error {

	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scanned value is incompatible (not a []byte): %#v", value)
	}

	return json.Unmarshal(bytes, target)
}

func ValueJsonColumn(value any) (driver.Value, error) {

	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal %#v: %v", value, err)
	}

	return bytes, nil
}
