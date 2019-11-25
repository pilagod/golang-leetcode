package tree

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BinaryNodeSuite struct {
	suite.Suite
}

func TestBinaryNode(t *testing.T) {
	suite.Run(t, &BinaryNodeSuite{})
}

func (s *BinaryNodeSuite) TestBuild() {
	root := BuildBinary([]interface{}{1, nil, 2, nil, nil, 3, 4})
	s.Equal(root.Val, 1)
	s.Nil(root.Left)
	s.Equal(root.Right.Val, 2)
	s.Equal(root.Right.Left.Val, 3)
	s.Equal(root.Right.Right.Val, 4)
}
