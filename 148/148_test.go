package main

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/pilagod/golang-leetcode/pkg/linkedlist"
)

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite
}

func (s *Suite) TestSortList() {
	head := linkedlist.New([]int{4, 2, 1, 3})
	ptr := sortList(head)
	for _, v := range []int{1, 2, 3, 4} {
		s.Equal(v, ptr.Val)
		ptr = ptr.Next
	}
}
