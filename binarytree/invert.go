package binarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	// Swap pointers of root.left and root.right
	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp
	return root
}
