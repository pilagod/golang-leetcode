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

func (s *Suite) TestEqualSubstring() {
	for _, test := range []struct {
		s        string
		t        string
		maxCost  int
		expected int
	}{
		{"abc", "ade", 0, 1},
		{"abc", "abc", 1, 3},
		{"abcd", "bcdf", 3, 3},
		{"abcd", "cdef", 3, 1},
		{"abcd", "acde", 0, 1},
		{"krrgw", "zjxss", 19, 2},
		{"ujteygggjwxnfl", "nstsenrzttikoy", 43, 5},
	} {
		got := equalSubstring(test.s, test.t, test.maxCost)
		s.Equal(test.expected, got, "%s to %s", test.s, test.t)
	}
}
