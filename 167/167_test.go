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

func (s *Suite) TestTwoSum() {
	for _, test := range []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{1, 2},
		},
		{
			[]int{-9, -7, -5, -3, -1, 1, 3, 9},
			-2,
			[]int{3, 7},
		},
	} {
		got := twoSum(test.nums, test.target)
		s.Equal(test.expected, got, "nums = %v, target = %d", test.nums, test.target)
	}
}
