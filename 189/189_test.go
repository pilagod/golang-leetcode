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

func (s *Suite) TestRotateArray() {
	for _, test := range []struct {
		nums     []int
		step     int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2, 3, 4, 5}, 18, []int{3, 4, 5, 1, 2}},
	} {
		rotate(test.nums, test.step)
		s.Equal(test.expected, test.nums)
	}
}
