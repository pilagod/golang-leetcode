package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(values []int) *ListNode {
	var head, ptr *ListNode
	for _, v := range values {
		n := &ListNode{
			Val: v,
		}
		if head == nil {
			head = n
		}
		if ptr != nil {
			ptr.Next = n
		}
		ptr = n
	}
	return head
}
