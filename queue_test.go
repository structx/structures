package structs_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/structx/structs"
)

type QueueSuite struct {
	suite.Suite
	q *structs.Queue[int64]
}

func (s *QueueSuite) SetupTest() {
	s.q = structs.MustNewQueueInt64[int64]()
}

func (s *QueueSuite) TestEnqueue() {

	assert := s.Assert()
	assert.NoError(s.q.Enqueue(1))
	assert.NoError(s.q.Enqueue(2))
	assert.NoError(s.q.Enqueue(3))

	assert.Equal(int64(3), s.q.Size())
}

func (s *QueueSuite) TestDequeue() {

	assert := s.Assert()

	assert.NoError(s.q.Enqueue(1))
	assert.NoError(s.q.Enqueue(2))
	assert.NoError(s.q.Enqueue(3))

	assert.Equal(int64(3), s.q.Size())

	i, err := s.q.Dequeue()
	assert.NoError(err)

	assert.Equal(1, i)
	assert.Equal(int64(2), s.q.Size())

	i, err = s.q.Dequeue()
	assert.NoError(err)

	assert.Equal(2, i)
	assert.Equal(int64(1), s.q.Size())

	i, err = s.q.Dequeue()
	assert.NoError(err)

	// expect queue size swap to index

	assert.Equal(3, i)
	assert.Equal(int64(3), s.q.Size())
}

func (s *QueueSuite) TestPeek() {

	assert := s.Assert()

	s.q.Enqueue(1)
	s.q.Enqueue(2)
	s.q.Enqueue(3)

	assert.Equal(int64(3), s.q.Size())

	i := s.q.Peek()
	assert.Equal(1, i)
	assert.Equal(int64(3), s.q.Size())
}

func TestQueueSuite(t *testing.T) {
	suite.Run(t, new(QueueSuite))
}
