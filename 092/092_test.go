package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite
}

func (s *Suite) TestReverse() {
	for _, test := range []struct {
		values   []int
		m        int
		n        int
		expected []int
	}{
		{
			nil,
			1,
			10,
			nil,
		},
		{
			[]int{1, 2, 3, 4, 5},
			2,
			4,
			[]int{1, 4, 3, 2, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			1,
			5,
			[]int{5, 4, 3, 2, 1},
		},
		{
			[]int{1, 2, 3, 4, 5},
			2,
			5,
			[]int{1, 5, 4, 3, 2},
		},
		{
			[]int{1, 2, 3, 4, 5},
			1,
			4,
			[]int{4, 3, 2, 1, 5},
		},
		{
			[]int{1, 2, 3},
			3,
			3,
			[]int{1, 2, 3},
		},
		{
			[]int{1, 2, 3, 4, 5},
			3,
			4,
			[]int{1, 2, 4, 3, 5},
		},
	} {
		got := reverseBetween(buildList(test.values), test.m, test.n)
		s.Equal(test.expected, traverseList(got))
	}
}

func buildList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{
		Val: values[0],
	}
	cur := head
	for _, value := range values[1:] {
		node := &ListNode{
			Val: value,
		}
		cur.Next = node
		cur = node
	}
	return head
}

func traverseList(head *ListNode) (result []int) {
	for cur := head; cur != nil; cur = cur.Next {
		result = append(result, cur.Val)
	}
	return
}
