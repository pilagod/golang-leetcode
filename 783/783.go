package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	prev, min := -1, -1
	minDiff(root, &prev, &min)
	return min
}

func minDiff(node *TreeNode, prev *int, min *int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		minDiff(node.Left, prev, min)
	}
	if *prev != -1 {
		diff := node.Val - *prev
		if *min == -1 || diff < *min {
			*min = diff
		}
	}
	if diff := node.Val - *prev; diff < *min {
		*min = diff
	}
	*prev = node.Val
	if node.Right != nil {
		minDiff(node.Right, prev, min)
	}
}
