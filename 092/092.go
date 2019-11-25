package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || m == n {
		return head
	}
	var start, end, rStart, rEnd *ListNode
	i, prev, cur := 1, new(ListNode), head
	for cur != nil {
		if i == m-1 {
			start = cur
		}
		if i == n+1 {
			end = cur
		}
		if i == m {
			rStart = cur
		}
		if i == n {
			rEnd = cur
		}
		if i > m && i <= n {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		} else {
			prev = cur
			cur = cur.Next
		}
		i++
	}
	if start != nil {
		start.Next = rEnd
	}
	if end != nil {
		rStart.Next = end
	} else {
		rStart.Next = nil
	}
	if start != nil {
		return head
	}
	return rEnd
}
