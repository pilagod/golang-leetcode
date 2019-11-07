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
	got := dailyTemperatures(nil)
	s.Len(got, 0)
}

func (s *Suite) TestOne() {
	got := dailyTemperatures([]int{1})
	s.Equal([]int{0}, got)
}

func (s *Suite) TestMultiple() {
	for _, test := range []struct {
		T        []int
		expected []int
	}{
		{
			[]int{73, 74, 75, 71, 69, 72, 76, 73},
			[]int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			[]int{1, 1, 1},
			[]int{0, 0, 0},
		},
		{
			[]int{75, 74, 73, 72, 71, 80},
			[]int{5, 4, 3, 2, 1, 0},
		},
	} {
		got := dailyTemperatures(test.T)
		s.Equal(test.expected, got, test.T)
	}

}
