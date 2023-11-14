package structs

import (
	"errors"
	"fmt"
	"sync/atomic"
)

// Queue queue
type Queue[T int8 | int64] struct {
	index, size *atomic.Int64
	store       Store[T]
}

// NewQueueInt8 returns new queue with int8 store
func NewQueueInt8[T int8](store Store[T]) *Queue[T] {
	return &Queue[T]{
		index: &atomic.Int64{},
		size:  &atomic.Int64{},
		store: store,
	}
}

// MustNewQueueInt8 returns new queue with int8 store
func MustNewQueueInt8[T int8]() *Queue[T] {
	return &Queue[T]{
		index: &atomic.Int64{},
		size:  &atomic.Int64{},
		store: newQueueStoreInt8[T](),
	}
}

// NewQueueInt64 returns new queue with int64 store
func NewQueueInt64[T int64](store Store[T]) *Queue[T] {
	return &Queue[T]{
		index: &atomic.Int64{},
		size:  &atomic.Int64{},
		store: store,
	}
}

// MustNewQueueInt64 returns new queue with int8 store
func MustNewQueueInt64[T int64]() *Queue[T] {
	return &Queue[T]{
		index: &atomic.Int64{},
		size:  &atomic.Int64{},
		store: newQueueStoreInt64[T](),
	}
}

type queueStore[T int8 | int64 | string] struct {
	data map[T]interface{}
}

func newQueueStoreInt8[T int8]() *queueStore[T] {
	return &queueStore[T]{
		data: make(map[T]interface{}),
	}
}

func newQueueStoreInt64[T int64]() *queueStore[T] {
	return &queueStore[T]{
		data: make(map[T]interface{}),
	}
}

func newQueueStoreString[T string]() *queueStore[T] {
	return &queueStore[T]{
		data: make(map[T]interface{}),
	}
}

// Get get item from store
func (qs *queueStore[T]) Get(key T) (interface{}, error) {

	v, ok := qs.data[key]
	if !ok {
		return nil, errors.New("item not found")
	}

	return v, nil
}

// Set set item to store
func (qs *queueStore[T]) Set(key T, value interface{}) error {
	qs.data[key] = value
	return nil
}

// Delete delete item from store
func (qs *queueStore[T]) Delete(key T) error {
	delete(qs.data, key)
	return nil
}

// Exists check if item exists in store
func (qs *queueStore[T]) Exists(key T) (bool, error) {

	_, ok := qs.data[key]
	if !ok {
		return false, nil
	}

	return true, nil
}

// Enqueue add item to queue
func (q *Queue[T]) Enqueue(task interface{}) error {

	// add item to queue
	Tpayload := T(q.size.Load())
	err := q.store.Set(Tpayload, task)
	if err != nil {
		return fmt.Errorf("failed to add item to queue: %w", err)
	}

	q.size.Add(1)

	return nil
}

// Dequeue remove item from queue
func (q *Queue[T]) Dequeue() (interface{}, error) {

	// get item from queue
	TPayload := T(q.index.Load())
	i, err := q.store.Get(TPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to get item from queue: %w", err)
	}

	// delete item from queue
	err = q.store.Delete(TPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to delete item from queue: %w", err)
	}

	// increment index
	q.index.Add(1)

	// if queue size is less than index, swap size with index
	if q.size.Load() < q.index.Load() {
		q.size.Swap(q.index.Load())
		return i, nil
	}

	// decrement queue size
	q.size.Add(-1)

	return i, nil
}

// Peek get item from queue without removing it
func (q *Queue[T]) Peek() interface{} {

	// get item from queue
	TPayload := T(q.index.Load())
	i, err := q.store.Get(TPayload)
	if err != nil {
		return nil
	}

	return i
}

// Size get queue size
func (q *Queue[T]) Size() int64 {
	return q.size.Load()
}
