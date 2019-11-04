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

func (s *Suite) TestLessThanTwoNums() {
	got := checkSubarraySum([]int{}, 1)
	s.Equal(false, got)
	got = checkSubarraySum([]int{1}, 1)
	s.Equal(false, got)
}

func (s *Suite) TestMoreThanTwoNums() {
	for _, test := range []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{23, 2, 4, 6, 7}, 6, true},
		{[]int{23, 2, 4, 6, 7}, 39, false},
		{[]int{4, 6, 8}, -6, true},
		{[]int{1, 2, 3}, 0, false},
		{[]int{0, 0}, 0, true},
	} {
		got := checkSubarraySum(test.nums, test.k)
		s.Equal(test.expected, got, test.nums)
	}
}
