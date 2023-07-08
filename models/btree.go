package btree

type BinaryNode struct {
	value int
	left  *BinaryNode
	right *BinaryNode
}

func CreateNode(value int) *BinaryNode {
	return &BinaryNode{value: value}
}

func InsertNode(value int, tree *BinaryNode) {
	if value < tree.value {
		if tree.left == nil {
			tree.left = &BinaryNode{value: value}
		} else {
			InsertNode(value, tree.left)
		}
	} else {
		if tree.right == nil {
			tree.right = &BinaryNode{value: value}
		} else {
			InsertNode(value, tree.right)
		}
	}
}

func PrintTree(tree *BinaryNode) {
	// Print the tree
}
