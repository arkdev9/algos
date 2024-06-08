package binarytree

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func traverse(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	leftDepth := traverse(root.Left, depth+1)
	rightDepth := traverse(root.Right, depth+1)
	return max(leftDepth, rightDepth)
}

func maxDepth(root *TreeNode) int {
	return traverse(root, 0)
}
