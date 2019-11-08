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

func (s *Suite) TestMultiple() {
	for _, test := range []struct {
		A        []int
		expected []int
	}{
		{
			[]int{-4, -1, 0, 3, 10},
			[]int{0, 1, 9, 16, 100},
		},
		{
			[]int{-7, -3, 2, 3, 11},
			[]int{4, 9, 9, 49, 121},
		},
	} {
		got := sortedSquares(test.A)
		s.Equal(test.expected, got, "%v", test.A)
	}
}
