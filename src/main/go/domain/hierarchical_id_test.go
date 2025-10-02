package domain

import (
	"database/sql"
	"testing"

	"github.com/benizzio/open-asset-allocator/infra/util"
	"github.com/lib/pq"
)

func TestHierarchicalIdValue_AllNonNil(t *testing.T) {
	var a = "a"
	var b = "b"
	var c = "c"

	var id = HierarchicalId{&a, &b, &c}
	actual, err := id.Value()
	if err != nil {
		t.Fatalf("Value() returned error: %v", err)
	}

	// Expected representation using pq.Array on []sql.NullString
	expectedArray := []sql.NullString{
		{String: "a", Valid: true},
		{String: "b", Valid: true},
		{String: "c", Valid: true},
	}
	expected, err := pq.Array(expectedArray).Value()
	if err != nil {
		t.Fatalf("pq.Array.Value() returned error: %v", err)
	}

	if util.ValueToString(actual) != util.ValueToString(expected) {
		t.Fatalf("mismatch\n actual: %q\nwant: %q", util.ValueToString(actual), util.ValueToString(expected))
	}
}

func TestHierarchicalIdValue_WithNilLevels(t *testing.T) {
	var a = "a"
	var c = "c"

	var id = HierarchicalId{&a, nil, &c}
	actual, err := id.Value()
	if err != nil {
		t.Fatalf("Value() returned error: %v", err)
	}

	expectedArray := []sql.NullString{
		{String: "a", Valid: true},
		{String: "", Valid: false},
		{String: "c", Valid: true},
	}
	expected, err := pq.Array(expectedArray).Value()
	if err != nil {
		t.Fatalf("pq.Array.Value() returned error: %v", err)
	}

	if util.ValueToString(actual) != util.ValueToString(expected) {
		t.Fatalf("mismatch with nils\n actual: %q\nwant: %q", util.ValueToString(actual), util.ValueToString(expected))
	}
}
