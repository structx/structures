package algs_test

import (
	"testing"

	"github.com/strcutx/structs/algs"
	"github.com/stretchr/testify/suite"
)

type ConsistentHashingWeightlessSuite[T uint64] struct {
	suite.Suite
	ring *algs.ConsistentHashRingWeightless[T]
}

func (s *ConsistentHashingWeightlessSuite[T]) SetupTest() {
	s.ring = algs.NewConsistentHashRingWeightless[T]()
}

func (s *ConsistentHashingWeightlessSuite[T]) TestAddPoint() {

	assert := s.Assert()

	err := s.ring.AddPoint("10.1.1.1")
	assert.NoError(err)

	err = s.ring.AddPoint("10.1.2.1")
	assert.NoError(err)

	err = s.ring.AddPoint("10.1.3.1")
	assert.NoError(err)
}

func TestConsistentHashingWeightlessSuite(t *testing.T) {
	suite.Run(t, new(ConsistentHashingWeightlessSuite[uint64]))
}
