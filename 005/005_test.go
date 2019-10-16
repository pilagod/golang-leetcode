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

func (s *Suite) TestStringWithOnlyOneChar() {
	for _, str := range []string{"a", "b", "c"} {
		got := longestPalindrome(str)
		s.Equal(str, got)
	}
}

func (s *Suite) TestStringWithSinglePalindrome() {
	for _, test := range []struct {
		str      string
		expected string
	}{
		{"aa", "aa"},
		{"aba", "aba"},
		{"abc", "a"},
		{"abcdedcba", "abcdedcba"},
		{"abacdefgh", "aba"},
		{"abacdfgdcaba", "aba"},
	} {
		got := longestPalindrome(test.str)
		s.Equal(test.expected, got)
	}
}

func (s *Suite) TestStringWithMultiplePalindrome() {
	for _, test := range []struct {
		str      string
		expected string
	}{
		{"abcaaaaaaa", "aaaaaaa"},
		{"abcaaaaaaacba", "abcaaaaaaacba"},
	} {
		got := longestPalindrome(test.str)
		s.Equal(test.expected, got)
	}
}
