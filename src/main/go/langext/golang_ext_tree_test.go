package langext

import (
	"reflect"
	"sort"
	"testing"
)

// TestNewMapTree_InitializesCorrectly tests that NewMapTree creates a node with correct initial state.
//
// Authored by: GitHub Copilot
func TestNewMapTree_InitializesCorrectly(t *testing.T) {

	var root = NewMapTree("root")

	if root.value != "root" {
		t.Fatalf("value = %q, expected %q", root.value, "root")
	}

	if root.depth != 0 {
		t.Fatalf("depth = %d, expected %d", root.depth, 0)
	}

	if root.children == nil {
		t.Fatal("children map should not be nil")
	}

	if len(root.children) != 0 {
		t.Fatalf("children count = %d, expected %d", len(root.children), 0)
	}
}

// TestNewMapTree_WithDifferentTypes tests NewMapTree with different comparable types.
//
// Authored by: GitHub Copilot
func TestNewMapTree_WithDifferentTypes(t *testing.T) {

	var intRoot = NewMapTree(42)

	if intRoot.value != 42 {
		t.Fatalf("int value = %d, expected %d", intRoot.value, 42)
	}
}

// TestAddChild_CreatesNewChild tests that AddChild creates a new child node.
//
// Authored by: GitHub Copilot
func TestAddChild_CreatesNewChild(t *testing.T) {

	var root = NewMapTree("root")

	var child = root.AddChild("child1")

	if child == nil {
		t.Fatal("AddChild returned nil")
	}

	if child.value != "child1" {
		t.Fatalf("child value = %q, expected %q", child.value, "child1")
	}

	if child.depth != 1 {
		t.Fatalf("child depth = %d, expected %d", child.depth, 1)
	}

	if len(root.children) != 1 {
		t.Fatalf("root children count = %d, expected %d", len(root.children), 1)
	}
}

// TestAddChild_ReturnsExistingChild tests that AddChild returns existing child when called with same value.
//
// Authored by: GitHub Copilot
func TestAddChild_ReturnsExistingChild(t *testing.T) {

	var root = NewMapTree("root")

	var child1 = root.AddChild("child")
	var child2 = root.AddChild("child")

	if child1 != child2 {
		t.Fatal("AddChild should return the same node for the same value")
	}

	if len(root.children) != 1 {
		t.Fatalf("root children count = %d, expected %d", len(root.children), 1)
	}
}

// TestAddChild_MultipleChildren tests adding multiple different children.
//
// Authored by: GitHub Copilot
func TestAddChild_MultipleChildren(t *testing.T) {

	var root = NewMapTree("root")

	root.AddChild("child1")
	root.AddChild("child2")
	root.AddChild("child3")

	if len(root.children) != 3 {
		t.Fatalf("root children count = %d, expected %d", len(root.children), 3)
	}
}

// TestAddChild_IncreasesDepth tests that depth increases with each level.
//
// Authored by: GitHub Copilot
func TestAddChild_IncreasesDepth(t *testing.T) {

	var root = NewMapTree("root")
	var level1 = root.AddChild("level1")
	var level2 = level1.AddChild("level2")
	var level3 = level2.AddChild("level3")

	if level1.depth != 1 {
		t.Fatalf("level1 depth = %d, expected %d", level1.depth, 1)
	}

	if level2.depth != 2 {
		t.Fatalf("level2 depth = %d, expected %d", level2.depth, 2)
	}

	if level3.depth != 3 {
		t.Fatalf("level3 depth = %d, expected %d", level3.depth, 3)
	}
}

// TestAddBranch_CreatesFullPath tests that AddBranch creates the entire path of nodes.
//
// Authored by: GitHub Copilot
func TestAddBranch_CreatesFullPath(t *testing.T) {

	var root = NewMapTree("root")
	var branch = []string{"a", "b", "c"}

	var lastNode = root.AddBranch(branch)

	if lastNode.value != "c" {
		t.Fatalf("last node value = %q, expected %q", lastNode.value, "c")
	}

	if lastNode.depth != 3 {
		t.Fatalf("last node depth = %d, expected %d", lastNode.depth, 3)
	}

	// Verify the path exists
	var nodeA = root.children["a"]
	if nodeA == nil {
		t.Fatal("node 'a' should exist")
	}

	var nodeB = nodeA.children["b"]
	if nodeB == nil {
		t.Fatal("node 'b' should exist")
	}

	var nodeC = nodeB.children["c"]
	if nodeC == nil {
		t.Fatal("node 'c' should exist")
	}
}

// TestAddBranch_ReusesExistingNodes tests that AddBranch reuses existing nodes in the path.
//
// Authored by: GitHub Copilot
func TestAddBranch_ReusesExistingNodes(t *testing.T) {

	var root = NewMapTree("root")

	root.AddBranch([]string{"a", "b", "c"})
	root.AddBranch([]string{"a", "b", "d"})

	// Only one 'a' and one 'b' should exist
	if len(root.children) != 1 {
		t.Fatalf("root children count = %d, expected %d", len(root.children), 1)
	}

	var nodeA = root.children["a"]
	if len(nodeA.children) != 1 {
		t.Fatalf("node 'a' children count = %d, expected %d", len(nodeA.children), 1)
	}

	var nodeB = nodeA.children["b"]
	if len(nodeB.children) != 2 {
		t.Fatalf("node 'b' children count = %d, expected %d", len(nodeB.children), 2)
	}
}

// TestAddBranch_EmptyBranch tests AddBranch with an empty slice.
//
// Authored by: GitHub Copilot
func TestAddBranch_EmptyBranch(t *testing.T) {

	var root = NewMapTree("root")

	var result = root.AddBranch([]string{})

	if result != root {
		t.Fatal("AddBranch with empty slice should return the root node")
	}

	if len(root.children) != 0 {
		t.Fatalf("root children count = %d, expected %d", len(root.children), 0)
	}
}

// TestAddBranchBreakingOnZeroValues_BasicBranch tests basic branch addition with no zero values.
//
// Authored by: GitHub Copilot
func TestAddBranchBreakingOnZeroValues_BasicBranch(t *testing.T) {

	var root = NewMapTree("root")
	var branch = []string{"a", "b", "c"}

	root.AddBranchBreakingOnZeroValues(branch)

	// Should add a -> b -> c (in order from root)
	var nodeA = root.children["a"]
	if nodeA == nil {
		t.Fatal("node 'a' should exist as first child")
	}

	var nodeB = nodeA.children["b"]
	if nodeB == nil {
		t.Fatal("node 'b' should exist as child of 'a'")
	}

	var nodeC = nodeB.children["c"]
	if nodeC == nil {
		t.Fatal("node 'c' should exist as child of 'b'")
	}
}

// TestAddBranchBreakingOnZeroValues_StopsAtZeroValue tests that iteration stops at zero value.
//
// Authored by: GitHub Copilot
func TestAddBranchBreakingOnZeroValues_StopsAtZeroValue(t *testing.T) {

	var root = NewMapTree("root")
	var branch = []string{"a", "", "c"}

	root.AddBranchBreakingOnZeroValues(branch)

	// Should add only 'a' (stops at "")
	var nodeA = root.children["a"]
	if nodeA == nil {
		t.Fatal("node 'a' should exist")
	}

	if len(nodeA.children) != 0 {
		t.Fatalf("node 'a' children count = %d, expected %d", len(nodeA.children), 0)
	}
}

// TestAddBranchBreakingOnZeroValues_ZeroValueAtStart tests branch with zero value at the start.
//
// Authored by: GitHub Copilot
func TestAddBranchBreakingOnZeroValues_ZeroValueAtStart(t *testing.T) {

	var root = NewMapTree("root")
	var branch = []string{"", "b", "c"}

	root.AddBranchBreakingOnZeroValues(branch)

	// Should not add anything (starts with zero value)
	if len(root.children) != 0 {
		t.Fatalf("root children count = %d, expected %d", len(root.children), 0)
	}
}

// TestAddBranchBreakingOnZeroValues_IntZeroValue tests with int type and zero value.
//
// Authored by: GitHub Copilot
func TestAddBranchBreakingOnZeroValues_IntZeroValue(t *testing.T) {

	var root = NewMapTree(0)
	var branch = []int{1, 0, 3}

	root.AddBranchBreakingOnZeroValues(branch)

	// Should add only 1 (stops at 0)
	var node1 = root.children[1]
	if node1 == nil {
		t.Fatal("node 1 should exist")
	}

	if len(node1.children) != 0 {
		t.Fatalf("node 1 children count = %d, expected %d", len(node1.children), 0)
	}
}

// TestAddInvertedBranchBreakingOnZeroValues_BasicInversion tests basic inverted branch addition.
//
// Authored by: GitHub Copilot
func TestAddInvertedBranchBreakingOnZeroValues_BasicInversion(t *testing.T) {

	var root = NewMapTree("root")
	var branch = []string{"a", "b", "c"}

	root.AddInvertedBranchBreakingOnZeroValues(branch)

	// Should add c -> b -> a (in order from root)
	var nodeC = root.children["c"]
	if nodeC == nil {
		t.Fatal("node 'c' should exist as first child")
	}

	var nodeB = nodeC.children["b"]
	if nodeB == nil {
		t.Fatal("node 'b' should exist as child of 'c'")
	}

	var nodeA = nodeB.children["a"]
	if nodeA == nil {
		t.Fatal("node 'a' should exist as child of 'b'")
	}
}

// TestAddInvertedBranchBreakingOnZeroValues_StopsAtZeroValue tests that iteration stops at zero value.
//
// Authored by: GitHub Copilot
func TestAddInvertedBranchBreakingOnZeroValues_StopsAtZeroValue(t *testing.T) {

	var root = NewMapTree("root")
	var branch = []string{"a", "", "c"}

	root.AddInvertedBranchBreakingOnZeroValues(branch)

	// Should add only 'c' (stops at "")
	var nodeC = root.children["c"]
	if nodeC == nil {
		t.Fatal("node 'c' should exist")
	}

	if len(nodeC.children) != 0 {
		t.Fatalf("node 'c' children count = %d, expected %d", len(nodeC.children), 0)
	}
}

// TestAddInvertedBranchBreakingOnZeroValues_ZeroValueAtEnd tests branch with zero value at the end.
//
// Authored by: GitHub Copilot
func TestAddInvertedBranchBreakingOnZeroValues_ZeroValueAtEnd(t *testing.T) {

	var root = NewMapTree("root")
	var branch = []string{"a", "b", ""}

	root.AddInvertedBranchBreakingOnZeroValues(branch)

	// Should not add anything (starts with zero value)
	if len(root.children) != 0 {
		t.Fatalf("root children count = %d, expected %d", len(root.children), 0)
	}
}

// TestAddInvertedBranchBreakingOnZeroValues_IntZeroValue tests with int type and zero value.
//
// Authored by: GitHub Copilot
func TestAddInvertedBranchBreakingOnZeroValues_IntZeroValue(t *testing.T) {

	var root = NewMapTree(0)
	var branch = []int{1, 0, 3}

	root.AddInvertedBranchBreakingOnZeroValues(branch)

	// Should add only 3 (stops at 0)
	var node3 = root.children[3]
	if node3 == nil {
		t.Fatal("node 3 should exist")
	}

	if len(node3.children) != 0 {
		t.Fatalf("node 3 children count = %d, expected %d", len(node3.children), 0)
	}
}

// TestExtractBranches_LeafNode tests ExtractBranches on a node with no children.
//
// Authored by: GitHub Copilot
func TestExtractBranches_LeafNode(t *testing.T) {

	var leaf = NewMapTree("leaf")

	var branches = leaf.ExtractBranches()

	if len(branches) != 1 {
		t.Fatalf("branches count = %d, expected %d", len(branches), 1)
	}

	var expected = []string{"leaf"}
	if !reflect.DeepEqual(branches[0], expected) {
		t.Fatalf("branch = %v, expected %v", branches[0], expected)
	}
}

// TestExtractBranches_SinglePath tests ExtractBranches on a linear tree.
//
// Authored by: GitHub Copilot
func TestExtractBranches_SinglePath(t *testing.T) {

	var root = NewMapTree("root")
	root.AddBranch([]string{"a", "b", "c"})

	var branches = root.ExtractBranches()

	if len(branches) != 1 {
		t.Fatalf("branches count = %d, expected %d", len(branches), 1)
	}

	var expected = []string{"root", "a", "b", "c"}
	if !reflect.DeepEqual(branches[0], expected) {
		t.Fatalf("branch = %v, expected %v", branches[0], expected)
	}
}

// TestExtractBranches_MultiplePaths tests ExtractBranches with multiple branches.
//
// Authored by: GitHub Copilot
func TestExtractBranches_MultiplePaths(t *testing.T) {

	var root = NewMapTree("root")
	root.AddBranch([]string{"a", "b"})
	root.AddBranch([]string{"a", "c", "e"})
	root.AddBranch([]string{"d"})

	var branches = root.ExtractBranches()

	if len(branches) != 3 {
		t.Fatalf("branches count = %d, expected %d", len(branches), 3)
	}

	// Sort branches for deterministic comparison (map iteration order is not guaranteed)
	var sortedBranches = sortBranches(branches)

	var expected = [][]string{
		{"root", "a", "b"},
		{"root", "a", "c", "e"},
		{"root", "d"},
	}

	if !reflect.DeepEqual(sortedBranches, expected) {
		t.Fatalf("branches = %v, expected %v", sortedBranches, expected)
	}
}

// TestExtractBranches_DeepTree tests ExtractBranches with a deeper tree structure.
//
// Authored by: GitHub Copilot
func TestExtractBranches_DeepTree(t *testing.T) {

	var root = NewMapTree("root")
	root.AddBranch([]string{"1", "2", "3", "4", "5"})

	var branches = root.ExtractBranches()

	if len(branches) != 1 {
		t.Fatalf("branches count = %d, expected %d", len(branches), 1)
	}

	var expected = []string{"root", "1", "2", "3", "4", "5"}
	if !reflect.DeepEqual(branches[0], expected) {
		t.Fatalf("branch = %v, expected %v", branches[0], expected)
	}
}

// TestExtractBranches_IntType tests ExtractBranches with int type.
//
// Authored by: GitHub Copilot
func TestExtractBranches_IntType(t *testing.T) {

	var root = NewMapTree(0)
	root.AddBranch([]int{1, 2, 3})

	var branches = root.ExtractBranches()

	if len(branches) != 1 {
		t.Fatalf("branches count = %d, expected %d", len(branches), 1)
	}

	var expected = []int{0, 1, 2, 3}
	if !reflect.DeepEqual(branches[0], expected) {
		t.Fatalf("branch = %v, expected %v", branches[0], expected)
	}
}

// TestMapTreeNode_ComplexScenario tests a complex tree manipulation scenario.
//
// Authored by: GitHub Copilot
func TestMapTreeNode_ComplexScenario(t *testing.T) {

	// Build a tree like:
	//        root
	//       /    \
	//      a      d
	//     / \      \
	//    b   c      e
	//         \
	//          f
	var root = NewMapTree("root")
	root.AddBranch([]string{"a", "b"})
	root.AddBranch([]string{"a", "c", "f"})
	root.AddBranch([]string{"d", "e"})

	var branches = root.ExtractBranches()

	if len(branches) != 3 {
		t.Fatalf("branches count = %d, expected %d", len(branches), 3)
	}

	var sortedBranches = sortBranches(branches)

	var expected = [][]string{
		{"root", "a", "b"},
		{"root", "a", "c", "f"},
		{"root", "d", "e"},
	}

	if !reflect.DeepEqual(sortedBranches, expected) {
		t.Fatalf("branches = %v, expected %v", sortedBranches, expected)
	}
}

// sortBranches sorts branches by converting them to strings for comparison.
// This ensures deterministic order since map iteration is not guaranteed.
//
// Authored by: GitHub Copilot
func sortBranches(branches [][]string) [][]string {

	var result = make([][]string, len(branches))
	copy(result, branches)

	sort.Slice(
		result, func(i, j int) bool {

			var minLen = len(result[i])
			if len(result[j]) < minLen {
				minLen = len(result[j])
			}

			for k := 0; k < minLen; k++ {
				if result[i][k] != result[j][k] {
					return result[i][k] < result[j][k]
				}
			}

			return len(result[i]) < len(result[j])
		},
	)

	return result
}
