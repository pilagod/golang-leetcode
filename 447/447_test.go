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

func (s *Suite) TestNoPoint() {
	got := numberOfBoomerangs(nil)
	s.Equal(0, got)
	got = numberOfBoomerangs(make([][]int, 0))
	s.Equal(0, got)
}

func (s *Suite) TestLessThanThreePoints() {
	got := numberOfBoomerangs([][]int{
		{0, 0},
		{1, 1},
	})
	s.Equal(0, got)
}

func (s *Suite) TestMoreThanThreePoints() {
	for _, test := range []struct {
		points   [][]int
		expected int
	}{
		{[][]int{{0, 0}, {1, 0}, {2, 0}}, 2},
		{[][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}}, 8},
		{[][]int{{0, 0}, {1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
	} {
		got := numberOfBoomerangs(test.points)
		s.Equal(test.expected, got)
	}
}
