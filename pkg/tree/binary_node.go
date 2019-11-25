package tree

// BinaryNode binary node
type BinaryNode struct {
	Val   int
	Left  *BinaryNode
	Right *BinaryNode
}

// BuildBinary builds binary tree by value shape
func BuildBinary(values []interface{}) *BinaryNode {
	root := buildBinaryNode(values, 0)
	return root
}

func buildBinaryNode(values []interface{}, index int) *BinaryNode {
	if index >= len(values) || values[index] == nil {
		return nil
	}
	node := &BinaryNode{
		Val: values[index].(int),
	}
	node.Left = buildBinaryNode(values, 2*index+1)
	node.Right = buildBinaryNode(values, 2*index+2)
	return node
}
