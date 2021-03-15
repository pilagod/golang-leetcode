package main

import (
	"github.com/pilagod/golang-leetcode/pkg/linkedlist"
)

type ListNode = linkedlist.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := split(head)
	l := sortList(head)
	r := sortList(mid)
	return merge(l, r)
}

func split(head *ListNode) *ListNode {
	var midPrev *ListNode
	ptr := head
	for ptr != nil && ptr.Next != nil {
		if midPrev == nil {
			midPrev = head
		} else {
			midPrev = midPrev.Next
		}
		ptr = ptr.Next.Next
	}
	mid := midPrev.Next
	midPrev.Next = nil
	return mid
}

func merge(l, r *ListNode) *ListNode {
	dummy := &ListNode{
		Val: 0,
	}
	lp, rp, ptr := l, r, dummy
	for lp != nil && rp != nil {
		if lp.Val <= rp.Val {
			ptr.Next = lp
			lp = lp.Next
		} else {
			ptr.Next = rp
			rp = rp.Next
		}
		ptr = ptr.Next
	}
	if lp != nil {
		ptr.Next = lp
	} else {
		ptr.Next = rp
	}
	return dummy.Next
}
