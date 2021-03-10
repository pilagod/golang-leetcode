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
		s      string
		length int
	}{
		{"", 0},
		{"a", 1},
		{"aa", 1},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"abba", 2},
	} {
		result := lengthOfLongestSubstring(t.s)
		s.Equal(t.length, result, t.s)
	}
}
