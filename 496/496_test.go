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

func (s *Suite) TestZeroNums1AndNums2() {
	got := nextGreaterElement(nil, nil)
	s.Len(got, 0)
}

func (s *Suite) TestZeroNums1() {
	got := nextGreaterElement(nil, []int{1, 2, 3})
	s.Len(got, 0)
}

func (s *Suite) TestZeroNums2() {
	got := nextGreaterElement([]int{1, 2, 3}, nil)
	s.Equal([]int{-1, -1, -1}, got)
}

func (s *Suite) TestValidNums1AndNums2() {
	for _, test := range []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			[]int{4, 1, 2},
			[]int{1, 3, 4, 2},
			[]int{-1, 3, -1},
		},
		{
			[]int{2, 4},
			[]int{1, 2, 3, 4},
			[]int{3, -1},
		},
	} {
		got := nextGreaterElement(test.nums1, test.nums2)
		s.Equal(test.expected, got, "nums1 = %v, nums2 = %v", test.nums1, test.nums2)
	}
}
