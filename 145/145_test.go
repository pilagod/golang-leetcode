package main

import (
	"testing"

	"github.com/pilagod/golang-leetcode/pkg/tree"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

func (s *Suite) TestPostOrder() {
	for _, test := range []struct {
		values   []interface{}
		expected []int
	}{
		{
			[]interface{}{1, nil, 2, nil, nil, 3},
			[]int{3, 2, 1},
		},
		{
			[]interface{}{1, 2, 3, 4, 5, 6},
			[]int{4, 5, 2, 6, 3, 1},
		},
	} {
		root := tree.BuildBinary(test.values)
		got := postorderTraversal(root)
		s.Equal(test.expected, got)
	}

}
