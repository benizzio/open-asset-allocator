package langext

import (
	"fmt"
	"testing"
)

// Test types used to validate fmt.Sprint() behavior with Stringer.
//
// Authored by: GitHub Copilot
type demoStringer int

// String implements fmt.Stringer for demoStringer.
//
// Authored by: GitHub Copilot
func (d demoStringer) String() string { return fmt.Sprintf("S#%d", int(d)) }

// TestCustomSlice_PrettyString_String tests PrettyString with string elements.
//
// Authored by: GitHub Copilot
func TestCustomSlice_PrettyString_String(t *testing.T) {
	var slice = CustomSlice[string]{"a", "b", "c"}
	var actual = slice.PrettyString()
	var expected = "a, b, c"
	if actual != expected {
		t.Fatalf("PrettyString() = %q, expected %q", actual, expected)
	}
}

// TestCustomSlice_PrettyString_Int tests PrettyString with int elements.
//
// Authored by: GitHub Copilot
func TestCustomSlice_PrettyString_Int(t *testing.T) {
	var slice = CustomSlice[int]{1, 2, 3}
	var actual = slice.PrettyString()
	var expected = "1, 2, 3"
	if actual != expected {
		t.Fatalf("PrettyString() = %q, expected %q", actual, expected)
	}
}

// TestCustomSlice_PrettyString_Empty tests PrettyString with an empty slice.
//
// Authored by: GitHub Copilot
func TestCustomSlice_PrettyString_Empty(t *testing.T) {
	var slice = CustomSlice[string]{}
	var actual = slice.PrettyString()
	var expected = ""
	if actual != expected {
		t.Fatalf("PrettyString() = %q, expected %q", actual, expected)
	}
}

// TestCustomSlice_PrettyString_Single tests PrettyString with a single element.
//
// Authored by: GitHub Copilot
func TestCustomSlice_PrettyString_Single(t *testing.T) {
	var slice = CustomSlice[string]{"only"}
	var actual = slice.PrettyString()
	var expected = "only"
	if actual != expected {
		t.Fatalf("PrettyString() = %q, expected %q", actual, expected)
	}
}

// TestCustomSlice_PrettyString_Stringer tests PrettyString with a Stringer element type.
//
// Authored by: GitHub Copilot
func TestCustomSlice_PrettyString_Stringer(t *testing.T) {
	var slice = CustomSlice[demoStringer]{1, 2, 3}
	var actual = slice.PrettyString()
	var expected = "S#1, S#2, S#3"
	if actual != expected {
		t.Fatalf("PrettyString() = %q, expected %q", actual, expected)
	}
}

// TestCustomSlice_PrettyString_Struct tests PrettyString with struct elements.
// It relies on Go's default struct formatting via fmt.Sprint.
//
// Authored by: GitHub Copilot
func TestCustomSlice_PrettyString_Struct(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}

	var slice = CustomSlice[person]{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 40},
	}

	var actual = slice.PrettyString()
	var expected = "{Alice 30}, {Bob 40}"
	if actual != expected {
		t.Fatalf("PrettyString() = %q, expected %q", actual, expected)
	}
}
