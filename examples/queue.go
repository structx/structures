package examples

import (
	"log"

	"github.com/structx/structs"
)

// ExampleQueue example queue
func ExampleQueue() {

	// create new queue
	q := structs.MustNewQueueInt64[int64]()

	// enqueue some items
	err := q.Enqueue(1)
	if err != nil {
		log.Fatalf("failed to enqueue: %v", err)
	}

	err = q.Enqueue(2)
	if err != nil {
		log.Fatalf("failed to enqueue: %v", err)
	}

	// dequeue item
	i, err := q.Dequeue()
	if err != nil {
		log.Fatalf("failed to dequeue: %v", err)
	}

	// expected output: 1
	log.Printf("dequeued: %d", i)

	// peek item
	i = q.Peek()

	// expected output: 2
	log.Printf("peeked: %d", i)
}
