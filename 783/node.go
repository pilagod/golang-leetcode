package main

// TreeNode node of binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func from(ns []interface{}) *TreeNode {
	if len(ns) == 0 {
		return nil
	}
	root := &TreeNode{Val: ns[0].(int)}
	for _, n := range ns[1:] {
		if n == nil {
			continue
		}
		insert(root, n.(int))
	}
	return root
}

func insert(node *TreeNode, num int) {
	if num > node.Val {
		if node.Right == nil {
			node.Right = &TreeNode{Val: num}
		} else {
			insert(node.Right, num)
		}
		return
	}
	if node.Left == nil {
		node.Left = &TreeNode{Val: num}
	} else {
		insert(node.Left, num)
	}
}
