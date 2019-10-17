package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite
}

func (s *Suite) TestStampLengthLargerThanTarget() {
	moves := movesToStamp("abc", "a")
	s.Nil(moves)
}

func (s *Suite) TestStampEqualToTarget() {
	moves := movesToStamp("abc", "abc")
	s.Equal([]int{0}, moves)
}

func (s *Suite) TestNoAvailableMove() {
	for _, test := range []struct {
		stamp  string
		target string
	}{
		{"abc", "defgh"},
		{"abc", "abbc"},
	} {
		moves := movesToStamp(test.stamp, test.target)
		s.Nil(moves)
	}

}

func (s *Suite) TestAvailableMove() {
	for _, test := range []struct {
		stamp  string
		target string
	}{
		{"abc", "ababc"},
		{"abc", "ababc"},
		{"abca", "aabcaca"},
		{"qr", "qrqqqrqrqrrqrqr"},
	} {
		moves := movesToStamp(test.stamp, test.target)
		s.assertTarget(moves, test.stamp, test.target)
	}
}

func (s *Suite) assertTarget(moves []int, stamp, target string) {
	got := strings.Repeat("?", len(target))
	for _, m := range moves {
		got = got[:m] + stamp + got[m+len(stamp):]
	}
	s.Equal(target, got)
}
