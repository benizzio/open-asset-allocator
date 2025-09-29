package domain

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"testing"

	"github.com/lib/pq"
)

// valueToString normalizes driver.Value to a comparable string.
// It handles string and []byte, and falls back to fmt.Sprintf otherwise.
//
// Authored by: GitHub Copilot
func valueToString(v driver.Value) string {
	switch x := v.(type) {
	case string:
		return x
	case []byte:
		return string(x)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func TestHierarchicalIdValue_AllNonNil(t *testing.T) {
	var a = "a"
	var b = "b"
	var c = "c"

	var id = HierarchicalId{&a, &b, &c}
	got, err := id.Value()
	if err != nil {
		t.Fatalf("Value() returned error: %v", err)
	}

	// Expected representation using pq.Array on []sql.NullString
	expectedArray := []sql.NullString{{String: "a", Valid: true}, {String: "b", Valid: true}, {String: "c", Valid: true}}
	expected, err := pq.Array(expectedArray).Value()
	if err != nil {
		t.Fatalf("pq.Array.Value() returned error: %v", err)
	}

	if valueToString(got) != valueToString(expected) {
		t.Fatalf("mismatch\n got: %q\nwant: %q", valueToString(got), valueToString(expected))
	}
}

func TestHierarchicalIdValue_WithNilLevels(t *testing.T) {
	var a = "a"
	var c = "c"

	var id = HierarchicalId{&a, nil, &c}
	got, err := id.Value()
	if err != nil {
		t.Fatalf("Value() returned error: %v", err)
	}

	expectedArray := []sql.NullString{{String: "a", Valid: true}, {String: "", Valid: false}, {String: "c", Valid: true}}
	expected, err := pq.Array(expectedArray).Value()
	if err != nil {
		t.Fatalf("pq.Array.Value() returned error: %v", err)
	}

	if valueToString(got) != valueToString(expected) {
		t.Fatalf("mismatch with nils\n got: %q\nwant: %q", valueToString(got), valueToString(expected))
	}
}
