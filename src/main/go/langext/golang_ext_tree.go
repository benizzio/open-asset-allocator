package langext

type TreeNode[T comparable] interface {
}

type MapTreeNode[T comparable] struct {
	value    T
	children map[T]*MapTreeNode[T]
	depth    int
}

func (node *MapTreeNode[T]) AddChild(childValue T) *MapTreeNode[T] {
	if _, exists := node.children[childValue]; !exists {
		childNode := &MapTreeNode[T]{value: childValue, children: make(map[T]*MapTreeNode[T]), depth: node.depth + 1}
		node.children[childValue] = childNode
	}
	return node.children[childValue]
}

func (node *MapTreeNode[T]) AddBranch(branch []T) *MapTreeNode[T] {
	currentNode := node
	for _, value := range branch {
		currentNode = currentNode.AddChild(value)
	}
	return currentNode
}

func (node *MapTreeNode[T]) AddInvertedBranchBreakingOnZeroValues(branch []T) {
	currentNode := node
	for i := len(branch) - 1; i >= 0; i-- {
		value := branch[i]
		if IsZeroValue(value) {
			break
		}
		currentNode = currentNode.AddChild(value)
	}
}

func (node *MapTreeNode[T]) ExtractBranches() [][]T {

	var branches [][]T

	currentStem := []T{node.value}
	if len(node.children) == 0 {
		branches = append(branches, currentStem)
		return branches
	}

	for _, child := range node.children {
		var childBranches = child.ExtractBranches()
		for _, childBranch := range childBranches {
			var branch = append(currentStem, childBranch...)
			branches = append(branches, branch)
		}
	}

	return branches
}

func NewMapTree[T comparable](value T) *MapTreeNode[T] {
	return &MapTreeNode[T]{value: value, children: make(map[T]*MapTreeNode[T]), depth: 0}
}
