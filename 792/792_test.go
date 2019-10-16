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

func (s *Suite) TestSingleCharWord() {
	got := numMatchingSubseq("abcde", []string{"a", "b", "c", "d", "e"})
	s.Equal(5, got)
}

func (s *Suite) TestMultipleCharsWord() {
	got := numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"})
	s.Equal(3, got)
}
