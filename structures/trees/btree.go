package trees

import "fmt"

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
	if tree == nil {
		return
	}
	// Print the tree inorder
	PrintTree(tree.left)
	fmt.Printf("%d ", tree.value)
	PrintTree(tree.right)
}

func btreeTest() {
	tree := CreateNode(10)
	InsertNode(30, tree)
	InsertNode(20, tree)
	InsertNode(60, tree)
	InsertNode(150, tree)
	InsertNode(15, tree)
	InsertNode(38, tree)

	PrintTree(tree)
}
