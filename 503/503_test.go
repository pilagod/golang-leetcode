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
	got := nextGreaterElements(nil)
	s.Len(got, 0)
}

func (s *Suite) TestOne() {
	got := nextGreaterElements([]int{1})
	s.Equal([]int{-1}, got)
}

func (s *Suite) TestMultiple() {
	for _, test := range []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{1, 2, 1},
			[]int{2, -1, 2},
		},
		{
			[]int{1, 1, 1},
			[]int{-1, -1, -1},
		},
		{
			[]int{5, 4, 3, 2, 1},
			[]int{-1, 5, 5, 5, 5},
		},
	} {
		got := nextGreaterElements(test.nums)
		s.Equal(test.expected, got, test.nums)
	}
}
