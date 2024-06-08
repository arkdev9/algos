package binarytree

import (
	"fmt"
	"strings"
)

// DFS traversal with a string output
func serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	return "^" + fmt.Sprint(root.Val) + serialize(root.Left) + serialize(root.Right)
}
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	s1 := serialize(root)
	s2 := serialize(subRoot)
	fmt.Println(s1)
	fmt.Println(s2)
	return strings.Contains(s1, s2)
}
