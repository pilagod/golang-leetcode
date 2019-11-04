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

func (s *Suite) TestNoCandidate() {
	got := combinationSum(nil, 1)
	s.Equal([][]int(nil), got)
}

func (s *Suite) TestOneCandidate() {
	for _, test := range []struct {
		candidate int
		target    int
		expected  [][]int
	}{
		{1, 1, [][]int{{1}}},
		{2, 3, [][]int(nil)},
	} {
		got := combinationSum([]int{test.candidate}, test.target)
		s.Equal(test.expected, got)
	}
}

func (s *Suite) TestMultipleCandidates() {
	for _, test := range []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			[]int{2, 3, 6, 7},
			7,
			[][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			[]int{2, 3, 5},
			8,
			[][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			[]int{2, 3, 4, 6, 8},
			8,
			[][]int{
				{2, 2, 2, 2},
				{2, 2, 4},
				{2, 3, 3},
				{2, 6},
				{4, 4},
				{8},
			},
		},
		{
			[]int{7, 3, 2},
			18,
			[][]int{
				{2, 2, 2, 2, 2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2, 2, 3, 3},
				{2, 2, 2, 2, 3, 7},
				{2, 2, 2, 3, 3, 3, 3},
				{2, 2, 7, 7},
				{2, 3, 3, 3, 7},
				{3, 3, 3, 3, 3, 3},
			},
		},
	} {
		got := combinationSum(test.candidates, test.target)
		s.Len(got, len(test.expected))
		for _, e := range test.expected {
			s.Contains(got, e)
		}
	}
}
