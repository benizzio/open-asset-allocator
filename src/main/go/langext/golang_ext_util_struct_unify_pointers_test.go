package langext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUnifyStructPointers tests the UnifyStructPointers function
//
// Authored by: GitHub Copilot
func TestUnifyStructPointers(t *testing.T) {

	// Test case 1: Basic string pointers
	t.Run(
		"basic string pointers", func(t *testing.T) {
			// Define a test struct
			type Person struct {
				Name *string
				Age  int
			}

			// Force creation of different pointer instances for the same string value
			var createString = func(s string) *string {
				var stringCopy = s + "" // Force a new allocation
				return &stringCopy
			}

			// Create test data
			var name1 = createString("John")
			var name2 = createString("John")  // Same value as name1, but different pointer
			var name3 = createString("Alice") // Different value

			var people = []Person{
				{Name: name1, Age: 30},
				{Name: name2, Age: 30},
				{Name: name3, Age: 25},
			}

			// Verify initial state: pointers to same value are different
			if name1 == name2 {
				// If compiler optimized the pointers to be the same, we can't properly test unification
				t.Skip("Test environment doesn't allow for distinct pointers with same value")
			}

			// Run the function
			UnifyStructPointers(people)

			// Verify that pointers to the same value have been unified
			assert.Same(t, people[0].Name, people[1].Name, "Pointers to equal values should be unified")

			// Verify that pointers to different values remain different
			assert.NotSame(t, people[0].Name, people[2].Name, "Pointers to different values should remain different")

			// Check the actual values are still correct
			assert.Equal(t, "John", *people[0].Name)
			assert.Equal(t, "John", *people[1].Name)
			assert.Equal(t, "Alice", *people[2].Name)
		},
	)

	// Test case 2: Multiple pointer types
	t.Run(
		"multiple pointer types", func(t *testing.T) {
			// Define a test struct with multiple pointer fields
			type Product struct {
				Name        *string
				Price       *float64
				IsAvailable *bool
				Quantity    *int
			}

			// Create helper functions to ensure distinct pointers
			var createString = func(s string) *string {
				var stringCopy = s + ""
				return &stringCopy
			}
			var createFloat = func(f float64) *float64 {
				var floatCopy = f
				return &floatCopy
			}
			var createBool = func(b bool) *bool {
				var boolCopy = b
				return &boolCopy
			}
			var createInt = func(i int) *int {
				var intCopy = i
				return &intCopy
			}

			// Create test data
			var products = []Product{
				{
					Name:        createString("Widget"),
					Price:       createFloat(9.99),
					IsAvailable: createBool(true),
					Quantity:    createInt(100),
				},
				{
					Name:        createString("Widget"), // Same name
					Price:       createFloat(9.99),      // Same price
					IsAvailable: createBool(true),       // Same availability
					Quantity:    createInt(50),          // Different quantity
				},
				{
					Name:        createString("Gadget"), // Different name
					Price:       createFloat(19.99),     // Different price
					IsAvailable: createBool(true),       // Same availability
					Quantity:    createInt(30),          // Different quantity
				},
			}

			// Run the function
			UnifyStructPointers(products)

			// Verify that pointers to the same values have been unified
			assert.Same(t, products[0].Name, products[1].Name, "Pointers to equal names should be unified")
			assert.Same(t, products[0].Price, products[1].Price, "Pointers to equal prices should be unified")
			assert.Same(
				t,
				products[0].IsAvailable,
				products[1].IsAvailable,
				"Pointers to equal availability should be unified",
			)
			assert.Same(
				t,
				products[0].IsAvailable,
				products[2].IsAvailable,
				"Pointers to equal availability should be unified",
			)

			// Verify that pointers to different values remain different
			assert.NotSame(
				t,
				products[0].Name,
				products[2].Name,
				"Pointers to different names should remain different",
			)
			assert.NotSame(
				t,
				products[0].Price,
				products[2].Price,
				"Pointers to different prices should remain different",
			)
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
		},
	)

	// Test case 3: Nested struct pointers
	t.Run(
		"nested struct pointers", func(t *testing.T) {
			// Define nested structs
			type Address struct {
				Street string
				City   string
			}

			type Person struct {
				Name    string
				Address *Address
			}

			// Create test data with explicitly different pointers
			var addr1 = &Address{Street: "123 Main St", City: "Anytown"}
			// Force creation of a different pointer with the same values
			var addr2 = new(Address)
			addr2.Street = addr1.Street
			addr2.City = addr1.City

			var addr3 = &Address{Street: "456 Oak St", City: "Othertown"}

			var people = []Person{
				{Name: "Person1", Address: addr1},
				{Name: "Person2", Address: addr2},
				{Name: "Person3", Address: addr3},
			}

			// Verify that the pointers are initially different
			assert.NotSame(t, addr1, addr2, "addr1 and addr2 should be different pointers")

			// Run the function
			UnifyStructPointers(people)

			// Verify that pointers to the same address have been unified
			assert.Same(t, people[0].Address, people[1].Address, "Pointers to equal addresses should be unified")

			// Verify that pointers to different addresses remain different
			assert.NotSame(
				t,
				people[0].Address,
				people[2].Address,
				"Pointers to different addresses should remain different",
			)

			// Check the actual values are still correct
			assert.Equal(t, "123 Main St", people[0].Address.Street)
			assert.Equal(t, "123 Main St", people[1].Address.Street)
			assert.Equal(t, "456 Oak St", people[2].Address.Street)
			assert.Equal(t, "Anytown", people[0].Address.City)
			assert.Equal(t, "Anytown", people[1].Address.City)
			assert.Equal(t, "Othertown", people[2].Address.City)
		},
	)

	// Test case 4: Empty slice
	t.Run(
		"empty slice", func(t *testing.T) {
			// Define a test struct
			type TestStruct struct {
				Name *string
			}

			// Create an empty slice
			var testData []TestStruct

			// This should not panic
			UnifyStructPointers(testData)

			// No assertion needed - we just verify it doesn't panic
		},
	)

	// Test case 5: Slice with single item
	t.Run(
		"single item slice", func(t *testing.T) {
			// Define a test struct
			type TestStruct struct {
				Name *string
			}

			// Create a slice with a single item
			var name = "test"
			var testData = []TestStruct{{Name: &name}}
			var originalPtr = testData[0].Name

			// Run the function
			UnifyStructPointers(testData)

			// Verify the pointer is unchanged
			assert.Same(t, originalPtr, testData[0].Name, "Pointer should remain unchanged for single item slice")
			assert.Equal(t, "test", *testData[0].Name, "Value should remain unchanged")
		},
	)

	// Test case 6: Mixed nil and non-nil pointers
	t.Run(
		"mixed nil and non-nil pointers", func(t *testing.T) {
			// Define a test struct
			type TestStruct struct {
				Name *string
			}

			// Create test data with some nil pointers
			var name1 = "test"
			var name2 = "test" + "" // Force different allocation

			var testData = []TestStruct{
				{Name: &name1},
				{Name: nil},
				{Name: &name2},
			}

			// Run the function
			UnifyStructPointers(testData)

			// Verify that nil pointers remain nil
			assert.Nil(t, testData[1].Name, "Nil pointers should remain nil")

			// Verify that pointers to the same value are unified
			assert.Same(t, testData[0].Name, testData[2].Name, "Pointers to equal values should be unified")

			// Check the actual values are still correct
			assert.Equal(t, "test", *testData[0].Name)
			assert.Equal(t, "test", *testData[2].Name)
		},
	)
}
