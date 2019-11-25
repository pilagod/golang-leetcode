package main

import (
	"github.com/pilagod/golang-leetcode/pkg/tree"
)

// TreeNode type alias for BinaryNode
type TreeNode = tree.BinaryNode

func postorderTraversal(root *TreeNode) []int {
	return postorder(root)
}

func postorder(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	values := postorder(node.Left)
	values = append(values, postorder(node.Right)...)
	values = append(values, node.Val)
	return values
}
