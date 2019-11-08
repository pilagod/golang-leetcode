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

func (s *Suite) TestLessThanThree() {
	got := threeSum([]int{1, 2})
	s.Len(got, 0)
}

func (s *Suite) TestThreeSum() {
	for _, test := range []struct {
		nums     []int
		expected [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, 0, 1}, {-1, -1, 2}},
		},
		{
			[]int{0, 0, 0, 0, 0},
			[][]int{{0, 0, 0}},
		},
	} {
		got := threeSum(test.nums)
		s.Len(got, len(test.expected))
		for _, e := range test.expected {
			s.Contains(got, e, test.nums)
		}
	}
}
