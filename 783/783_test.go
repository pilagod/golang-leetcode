package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

func (s *Suite) TestBuildTreeFromArray() {
	nums := []interface{}{4, 2, 6, 1, 3, nil, nil}
	root := from(nums)
	s.Equal(4, root.Val)
	s.Equal(2, root.Left.Val)
	s.Equal(6, root.Right.Val)
	s.Equal(1, root.Left.Left.Val)
	s.Equal(3, root.Left.Right.Val)
	s.Equal((*TreeNode)(nil), root.Right.Left)
	s.Equal((*TreeNode)(nil), root.Right.Right)
}

func (s *Suite) TestMinDiffBST() {
	for _, test := range []struct {
		root     *TreeNode
		expected int
	}{
		{from([]interface{}{4, 2, 6}), 2},
		{from([]interface{}{4, 2, 6, 1, 3, nil, nil}), 1},
		{from([]interface{}{90, 69, nil, 49, 89, nil, 52, nil, nil, nil, nil}), 1},
	} {
		got := minDiffInBST(test.root)
		s.Equal(test.expected, got)
	}
}
