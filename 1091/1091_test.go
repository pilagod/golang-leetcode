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

func (s *Suite) TestLengthOfLongestSubstring() {
	for _, t := range []struct {
		grid   [][]int
		length int
	}{
		{
			[][]int{
				{0},
			},
			1,
		},
		{
			[][]int{
				{0, 1},
				{1, 0},
			},
			2,
		},
		{
			[][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			4,
		},
		{
			[][]int{
				{1, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			-1,
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 1},
				{0, 1, 0},
			},
			-1,
		},
		{
			[][]int{
				{0, 1, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 0},
				{0, 1, 1, 0, 1, 0},
				{0, 0, 0, 0, 1, 0},
				{1, 1, 1, 1, 1, 0},
				{1, 1, 1, 1, 1, 0},
			},
			14,
		},
	} {
		result := shortestPathBinaryMatrix(t.grid)
		s.Equal(t.length, result, t.grid)
	}
}
