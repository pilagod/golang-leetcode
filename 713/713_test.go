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

func (s *Suite) TestEmpty() {
	got := numSubarrayProductLessThanK(nil, 1)
	s.Equal(0, got)
}

func (s *Suite) TestMultiple() {
	for _, test := range []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{10, 5, 2, 6}, 100, 8},
		{[]int{1, 2, 3}, 100, 6},
		{[]int{1, 1, 1}, 0, 0},
		{[]int{8, 10, 1, 3}, 1, 0},
	} {
		got := numSubarrayProductLessThanK(test.nums, test.k)
		s.Equal(test.expected, got, "nums = %v, k = %d", test.nums, test.k)
	}
}
