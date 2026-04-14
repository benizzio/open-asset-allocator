package domain

import (
	"database/sql"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/benizzio/open-asset-allocator/infra/util"
)

func TestHierarchicalIdValue_AllNonNil(t *testing.T) {
	var a = "a"
	var b = "b"
	var c = "c"

	var id = HierarchicalId{&a, &b, &c}
	actual, err := id.Value()
	require.NoError(t, err)

	// Expected representation using pq.Array on []sql.NullString
	expectedArray := []sql.NullString{
		{String: "a", Valid: true},
		{String: "b", Valid: true},
		{String: "c", Valid: true},
	}
	expected, err := pq.Array(expectedArray).Value()
	require.NoError(t, err)

	assert.Equal(t, util.ValueToString(expected), util.ValueToString(actual))
}

func TestHierarchicalIdValue_WithNilLevels(t *testing.T) {
	var a = "a"
	var c = "c"

	var id = HierarchicalId{&a, nil, &c}
	actual, err := id.Value()
	require.NoError(t, err)

	expectedArray := []sql.NullString{
		{String: "a", Valid: true},
		{String: "", Valid: false},
		{String: "c", Valid: true},
	}
	expected, err := pq.Array(expectedArray).Value()
	require.NoError(t, err)

	assert.Equal(t, util.ValueToString(expected), util.ValueToString(actual))
}

func TestHierarchicalIdIsTopLevel_Empty(t *testing.T) {
	var hierarchicalId = HierarchicalId{}

	assert.False(t, hierarchicalId.IsTopLevel())
}

func TestHierarchicalIdParentLevelId_Empty(t *testing.T) {
	var hierarchicalId = HierarchicalId{}

	var parentLevelId = hierarchicalId.ParentLevelId()
	assert.Empty(t, parentLevelId)
}
