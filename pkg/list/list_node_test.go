package list

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

func (s *Suite) TestNew() {
	values := []int{3, 5, 4}
	i := 0
	for head := New(values); head != nil; head = head.Next {
		s.Equal(values[i], head.Val, i)
		i++
	}
	s.Equal(3, i)
}

func (s *Suite) TestNewWithEmptyValues() {
	s.Nil(New(nil))
}
