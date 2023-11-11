// Package algs algorithms package using data structures
package algs

import (
	"fmt"
	"hash/fnv"

	"github.com/strcutx/structs"
)

// ConsistentHashRingWeightless consistent hashing ring with no weights
type ConsistentHashRingWeightless[T uint64] struct {
	tree *structs.BST[T]   // tree used to quickly lookup angles
	ring map[uint64]string // store server angles on ring
}

// NewConsistentHashRingWeightless create new consistent hashing ring
func NewConsistentHashRingWeightless[T uint64]() *ConsistentHashRingWeightless[T] {
	return &ConsistentHashRingWeightless[T]{
		tree: structs.NewBST[T](),
		ring: make(map[uint64]string),
	}
}

// AddPoint add point to ring
func (cw *ConsistentHashRingWeightless[T]) AddPoint(value string) error {

	// convert label to uint64 hash
	fh := fnv.New64()
	_, err := fh.Write([]byte(value))
	if err != nil {
		return fmt.Errorf("failed to hash value: %v", err)
	}

	h := fh.Sum64()

	// module hash by 360
	// ensure value is less than 360 degress
	a := h % 360

	// add angle as new node in tree
	cw.tree.AddNodeUint64(a)

	// add angle and label to ring
	cw.ring[a] = value

	return nil
}

// Distribute distribute data to server
func (cw *ConsistentHashRingWeightless[T]) Distribute(data []byte) (string, error) {

	// convert data to uint64 hash
	fh := fnv.New64()
	_, err := fh.Write(data)
	if err != nil {
		return "", fmt.Errorf("failed to hash data: %v", err)
	}

	h := fh.Sum64()

	// module hash by 360
	// ensure value is less than 360 degress
	a := h % 360

	// find closest angle in tree
	ca, found := cw.tree.FindClosestUint64(a)
	if !found {
		return "", &structs.ErrNotFound[T]{Payload: T(a)}
	}

	// return value for closest angle
	return cw.ring[ca], nil
}

// DeleteAll delete all points from ring
func (cw *ConsistentHashRingWeightless[T]) DeleteAll(angle uint64) error {

	// find ip for angle
	ip := cw.ring[angle]
	if ip == "" {
		Tpayload := T(angle)
		return &structs.ErrNotFound[T]{Payload: Tpayload}
	}

	// delete angle from tree and
	// delete angle from ring
	for k, v := range cw.ring {
		if v == ip {
			cw.tree.DeleteUint64(k)
			delete(cw.ring, k)
		}
	}

	return nil
}

// DeletePoint delete point from ring
func (cw *ConsistentHashRingWeightless[T]) DeletePoint(angle T) error {

	// delete angle from tree
	cw.tree.DeleteUint64(uint64(angle))

	// delete angle from ring
	delete(cw.ring, uint64(angle))

	return nil
}
