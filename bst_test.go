package structs_test

import (
	"log"
	"testing"

	"github.com/strcutx/structs"
	"github.com/stretchr/testify/suite"
)

type BSTSuite struct {
	suite.Suite
	tree *structs.BST[int64]
}

func (s *BSTSuite) SetupTest() {
	s.tree = structs.NewBST[int64]()
}

func (s *BSTSuite) TestAddNode() {

	s.tree.AddNodeInt64(1)
	s.tree.AddNodeInt64(2)
	s.tree.AddNodeInt64(3)
	s.tree.AddNodeInt64(5)
	s.tree.AddNodeInt64(7)
	s.tree.AddNodeInt64(9)
	s.tree.AddNodeInt64(8)
	s.tree.AddNodeInt64(13)
	s.tree.AddNodeInt64(14)

	log.Println(s.tree.Stringify())
}

func (s *BSTSuite) TestFindClosest() {

	assert := s.Assert()

	s.tree.AddNodeInt64(1)
	s.tree.AddNodeInt64(2)
	s.tree.AddNodeInt64(3)
	s.tree.AddNodeInt64(5)
	s.tree.AddNodeInt64(7)
	s.tree.AddNodeInt64(9)
	s.tree.AddNodeInt64(8)
	s.tree.AddNodeInt64(13)
	s.tree.AddNodeInt64(14)
	s.tree.AddNodeInt64(4)
	s.tree.AddNodeInt64(22)
	s.tree.AddNodeInt64(21)
	s.tree.AddNodeInt64(100)
	s.tree.AddNodeInt64(101)

	payload, found := s.tree.FindClosestUint64(85)
	assert.Equal(true, found)

	log.Printf("lowest node payload: %d", payload)
}

func (s *BSTSuite) TestDelete() {

	s.tree.AddNodeInt64(1)
	s.tree.AddNodeInt64(2)
	s.tree.AddNodeInt64(3)
	s.tree.AddNodeInt64(5)
	s.tree.AddNodeInt64(7)
	s.tree.AddNodeInt64(9)
	s.tree.AddNodeInt64(8)
	s.tree.AddNodeInt64(13)
	s.tree.AddNodeInt64(14)
	s.tree.AddNodeInt64(4)
	s.tree.AddNodeInt64(22)
	s.tree.AddNodeInt64(21)
	s.tree.AddNodeInt64(100)
	s.tree.AddNodeInt64(101)

	log.Println(s.tree.Stringify())

	s.tree.DeleteInt64(10000)
}

func TestBSTSuite(t *testing.T) {
	suite.Run(t, new(BSTSuite))
}
