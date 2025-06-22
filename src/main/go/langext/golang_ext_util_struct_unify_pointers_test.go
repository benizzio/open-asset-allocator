package langext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// createStringPointer returns a pointer to a new string copy.
//
// Authored by: GitHub Copilot
func createStringPointer(s string) *string {

	var stringCopy = s + ""
	return &stringCopy
}

// createFloatPointer returns a pointer to a new float64 copy.
//
// Authored by: GitHub Copilot
func createFloatPointer(f float64) *float64 {

	var floatCopy = f
	return &floatCopy
}

// createBoolPointer returns a pointer to a new bool copy.
//
// Authored by: GitHub Copilot
func createBoolPointer(b bool) *bool {

	var boolCopy = b
	return &boolCopy
}

// createIntPointer returns a pointer to a new int copy.
//
// Authored by: GitHub Copilot
func createIntPointer(i int) *int {

	var intCopy = i
	return &intCopy
}

// TestUnifyStructPointersBasicStringPointers tests the unification of pointers to string values.
//
// This test verifies that pointers to equal string values are unified, while pointers
// to different string values remain distinct.
//
// Authored by: GitHub Copilot
func TestUnifyStructPointersBasicStringPointers(t *testing.T) {

	// Define a test struct
	type Person struct {
		Name *string
		Age  int
	}

	// Create test data with explicitly different pointers to the same value
	var name1 = createStringPointer("John")
	var name2 = createStringPointer("John")  // Same value as name1, but different pointer
	var name3 = createStringPointer("Alice") // Different value

	var people = []Person{
		{Name: name1, Age: 30},
		{Name: name2, Age: 30},
		{Name: name3, Age: 25},
	}

	// Verify initial condition: pointers to the same value are different instances
	assert.NotSame(t, people[0].Name, people[1].Name, "Pointers to same value should not be the same initially")

	// Run the function under test
	UnifyStructPointers(people)

	// Verify that pointers to the same value have been unified
	assert.Same(t, people[0].Name, people[1].Name, "Pointers to equal values should be unified")

	// Verify that pointers to different values remain different
	assert.NotSame(t, people[0].Name, people[2].Name, "Pointers to different values should remain different")

	// Check the actual values are still correct
	assert.Equal(t, "John", *people[0].Name)
	assert.Equal(t, "John", *people[1].Name)
	assert.Equal(t, "Alice", *people[2].Name)
}

// TestUnifyStructPointersMultiplePointerTypes tests unification for structs with multiple pointer field types.
//
// This test ensures that pointers to equal values are unified and pointers to different values
// remain distinct across different primitive types (string, float64, bool, int).
//
// Authored by: GitHub Copilot
func TestUnifyStructPointersMultiplePointerTypes(t *testing.T) {

	// Product represents a product with multiple pointer fields for testing pointer unification.
	type Product struct {
		Name        *string
		Price       *float64
		IsAvailable *bool
		Quantity    *int
	}

	// Create test data with explicitly different pointers to equal values
	var products = []Product{
		{
			Name:        createStringPointer("Widget"),
			Price:       createFloatPointer(9.99),
			IsAvailable: createBoolPointer(true),
			Quantity:    createIntPointer(100),
		},
		{
			Name:        createStringPointer("Widget"), // Same name
			Price:       createFloatPointer(9.99),      // Same price
			IsAvailable: createBoolPointer(true),       // Same availability
			Quantity:    createIntPointer(50),          // Different quantity
		},
		{
			Name:        createStringPointer("Gadget"), // Different name
			Price:       createFloatPointer(19.99),     // Different price
			IsAvailable: createBoolPointer(true),       // Same availability
			Quantity:    createIntPointer(30),          // Different quantity
		},
	}

	// Verify initial conditions: pointers to equal values are different instances
	assert.NotSame(t, products[0].Name, products[1].Name, "Pointers to equal names should be different initially")
	assert.NotSame(t, products[0].Price, products[1].Price, "Pointers to equal prices should be different initially")
	assert.NotSame(
		t,
		products[0].IsAvailable,
		products[1].IsAvailable,
		"Pointers to equal availability should be different initially",
	)
	assert.NotSame(
		t,
		products[0].IsAvailable,
		products[2].IsAvailable,
		"Pointers to equal availability should be different initially",
	)

	// Run the function under test
	UnifyStructPointers(products)

	// Verify that pointers to equal values have been unified
	assert.Same(t, products[0].Name, products[1].Name, "Pointers to equal names should be unified")
	assert.Same(t, products[0].Price, products[1].Price, "Pointers to equal prices should be unified")
	assert.Same(t, products[0].IsAvailable, products[1].IsAvailable, "Pointers to equal availability should be unified")
	assert.Same(t, products[0].IsAvailable, products[2].IsAvailable, "Pointers to equal availability should be unified")

	// Verify that pointers to different values remain different
	assert.NotSame(t, products[0].Name, products[2].Name, "Pointers to different names should remain different")
	assert.NotSame(t, products[0].Price, products[2].Price, "Pointers to different prices should remain different")
	assert.NotSame(
		t,
		products[0].Quantity,
		products[1].Quantity,
		"Pointers to different quantities should remain different",
	)
	assert.NotSame(
		t,
		products[0].Quantity,
		products[2].Quantity,
		"Pointers to different quantities should remain different",
	)

	// Check the actual values are still correct
	assert.Equal(t, "Widget", *products[0].Name)
	assert.Equal(t, "Widget", *products[1].Name)
	assert.Equal(t, "Gadget", *products[2].Name)
	assert.Equal(t, 9.99, *products[0].Price)
	assert.Equal(t, 9.99, *products[1].Price)
	assert.Equal(t, 19.99, *products[2].Price)
	assert.Equal(t, true, *products[0].IsAvailable)
	assert.Equal(t, true, *products[1].IsAvailable)
	assert.Equal(t, true, *products[2].IsAvailable)
	assert.Equal(t, 100, *products[0].Quantity)
	assert.Equal(t, 50, *products[1].Quantity)
	assert.Equal(t, 30, *products[2].Quantity)
}

// TestUnifyStructPointersNestedStructPointers tests unification for nested struct pointer fields.
//
// This test verifies that pointers to equal struct values are unified,
// while pointers to different struct values remain distinct.
//
// Authored by: GitHub Copilot
func TestUnifyStructPointersNestedStructPointers(t *testing.T) {

	// Define nested structs for testing
	type Address struct {
		Street string
		City   string
	}

	type Person struct {
		Name    string
		Address *Address
	}

	// Create test data with explicitly different pointers to equal struct values
	var addr1 = &Address{Street: "123 Main St", City: "Anytown"}
	var addr2 = new(Address)
	addr2.Street = addr1.Street
	addr2.City = addr1.City
	var addr3 = &Address{Street: "456 Oak St", City: "Othertown"}

	var people = []Person{
		{Name: "Person1", Address: addr1},
		{Name: "Person2", Address: addr2},
		{Name: "Person3", Address: addr3},
	}

	// Verify initial condition: pointers to the same struct value are different instances
	assert.NotSame(t, addr1, addr2, "addr1 and addr2 should be different pointers")

	// Run the function under test
	UnifyStructPointers(people)

	// Verify that pointers to equal struct values have been unified
	assert.Same(t, people[0].Address, people[1].Address, "Pointers to equal addresses should be unified")

	// Verify that pointers to different struct values remain different
	assert.NotSame(t, people[0].Address, people[2].Address, "Pointers to different addresses should remain different")

	// Check the actual values are still correct
	assert.Equal(t, "123 Main St", people[0].Address.Street)
	assert.Equal(t, "123 Main St", people[1].Address.Street)
	assert.Equal(t, "456 Oak St", people[2].Address.Street)
	assert.Equal(t, "Anytown", people[0].Address.City)
	assert.Equal(t, "Anytown", people[1].Address.City)
	assert.Equal(t, "Othertown", people[2].Address.City)
}

// TestUnifyStructPointersEmptySlice tests that an empty slice does not cause any issues.
//
// This test verifies that the function handles empty slices gracefully without panicking.
//
// Authored by: GitHub Copilot
func TestUnifyStructPointersEmptySlice(t *testing.T) {

	// Define a test struct
	type TestStruct struct {
		Name *string
	}

	// Create an empty slice
	var testData []TestStruct

	// Verify the function handles empty slices without panicking
	UnifyStructPointers(testData)
}

// TestUnifyStructPointersSingleItemSlice tests that a single-item slice is unchanged.
//
// This test verifies that the function does not modify pointers when there's only one item.
//
// Authored by: GitHub Copilot
func TestUnifyStructPointersSingleItemSlice(t *testing.T) {

	// Define a test struct
	type TestStruct struct {
		Name *string
	}

	// Create a slice with a single item
	var name = "test"
	var testData = []TestStruct{{Name: &name}}
	var originalPtr = testData[0].Name

	// Run the function under test
	UnifyStructPointers(testData)

	// Verify the pointer is unchanged
	assert.Same(t, originalPtr, testData[0].Name, "Pointer should remain unchanged for single item slice")
	assert.Equal(t, "test", *testData[0].Name, "Value should remain unchanged")
}

// TestUnifyStructPointersMixedNilAndNonNil tests handling of nil pointers.
//
// This test verifies that nil pointers remain nil and non-nil pointers to equal values are unified.
//
// Authored by: GitHub Copilot
func TestUnifyStructPointersMixedNilAndNonNil(t *testing.T) {

	// Define a test struct
	type TestStruct struct {
		Name *string
	}

	// Create test data with a mix of nil and non-nil pointers
	var name1 = "test"
	var name2 = "test" + "" // Force different allocation
	var testData = []TestStruct{
		{Name: &name1},
		{Name: nil},
		{Name: &name2},
	}

	// Run the function under test
	UnifyStructPointers(testData)

	// Verify that nil pointers remain nil
	assert.Nil(t, testData[1].Name, "Nil pointers should remain nil")

	// Verify that non-nil pointers to equal values are unified
	assert.Same(t, testData[0].Name, testData[2].Name, "Pointers to equal values should be unified")
	assert.Equal(t, "test", *testData[0].Name)
	assert.Equal(t, "test", *testData[2].Name)
}
