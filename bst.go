// Package structs structures package
package structs

import "fmt"

type bstNode[T int | int64 | int32 | uint | uint8 | uint32 | uint64] struct {
	left, right *bstNode[T]
	payload     T
}

// BST binary search tree
type BST[T int | int64 | int32 | uint | uint8 | uint32 | uint64] struct {
	head *bstNode[T]
}

// NewBST create new binary search tree
func NewBST[T int | int64 | int32 | uint | uint8 | uint32 | uint64]() *BST[T] {
	return &BST[T]{
		head: nil, // intentionally set nil
	}
}

// AddNodeUint add node to tree with uint payload
func (b *BST[T]) AddNodeUint(payload uint) {
	Tpayload := T(payload)
	b.insert(Tpayload)
}

// AddNodeUint8 add node to tree with uint8 payload
func (b *BST[T]) AddNodeUint8(payload uint8) {
	Tpayload := T(payload)
	b.insert(Tpayload)
}

// AddNodeUint32 add node to tree with uint32 payload
func (b *BST[T]) AddNodeUint32(payload uint32) {
	Tpayload := T(payload)
	b.insert(Tpayload)
}

// AddNodeUint64 add node to tree with uint64 payload
func (b *BST[T]) AddNodeUint64(payload uint64) {
	Tpayload := T(payload)
	b.insert(Tpayload)
}

// AddNodeInt add node to tree with int payload
func (b *BST[T]) AddNodeInt(payload int) {
	Tpayload := T(payload)
	b.insert(Tpayload)
}

// AddNodeInt32 add node to tree with int32 payload
func (b *BST[T]) AddNodeInt32(payload int32) {
	Tpayload := T(payload)
	b.insert(Tpayload)
}

// AddNodeInt64 add node to tree with int64 payload
func (b *BST[T]) AddNodeInt64(payload int64) {
	Tpayload := T(payload)
	b.insert(Tpayload)
}

func (b *BST[T]) insert(payload T) {

	// if empty tree add head node
	if b.head == nil {
		b.head = &bstNode[T]{payload: payload, left: nil, right: nil}
		return
	}

	// if payload is less than head node
	if payload < b.head.payload {

		if b.head.left == nil {
			b.head.left = &bstNode[T]{payload: payload, left: nil, right: nil}
			return
		}

		b.head.left.insert(payload)
		return
	}

	// if payload is greater than head node
	if b.head.right == nil {
		b.head.right = &bstNode[T]{payload: payload, left: nil, right: nil}
		return
	}

	b.head.right.insert(payload)
}

func (n *bstNode[T]) insert(payload T) {

	// if node is empty add node
	if n == nil {
		n = &bstNode[T]{payload: payload, left: nil, right: nil}
		return
	}

	// if payload is less than node
	if payload < n.payload {

		if n.left == nil {
			n.left = &bstNode[T]{payload: payload, left: nil, right: nil}
			return
		}

		n.left.insert(payload)
		return
	}

	// if payload is greater than node
	if n.right == nil {
		n.right = &bstNode[T]{payload: payload, left: nil, right: nil}
		return
	}

	n.right.insert(payload)
}

// DeleteUint delete node from tree with uint payload
func (b *BST[T]) DeleteUint(payload uint) {
	Tpayload := T(payload)
	b.head.delete(Tpayload)
}

// DeleteUint8 delete node from tree with uint8 payload
func (b *BST[T]) DeleteUint8(payload uint8) {
	Tpayload := T(payload)
	b.head.delete(Tpayload)
}

// DeleteUint32 delete node from tree with uint32 payload
func (b *BST[T]) DeleteUint32(payload uint32) {
	Tpayload := T(payload)
	b.head.delete(Tpayload)
}

// DeleteUint64 delete node from tree with uint64 payload
func (b *BST[T]) DeleteUint64(payload uint64) {
	Tpayload := T(payload)
	b.head.delete(Tpayload)
}

// DeleteInt delete node from tree with int payload
func (b *BST[T]) DeleteInt(payload int) {
	Tpayload := T(payload)
	b.head.delete(Tpayload)
}

// DeleteInt32 delete node from tree with int32 payload
func (b *BST[T]) DeleteInt32(payload int32) {
	Tpayload := T(payload)
	b.head.delete(Tpayload)
}

// DeleteInt64 delete node from tree with int64 payload
func (b *BST[T]) DeleteInt64(payload int64) {
	Tpayload := T(payload)
	b.head.delete(Tpayload)
}

func (b *BST[T]) delete(payload T) {
	if b.head == nil {
		return
	}

	if b.head.payload == payload {

		n := b.head

		// if both left and right are empty delete node
		if n.left == nil && n.right == nil {
			n = nil
		} else if n.left != nil && n.right == nil {
			// if only left right node is empty
			// move left node up to current node
			tmp := n.left
			n = nil
			n = tmp
			return
		}

		// left and right have nodes
		n = n.left
		return
	}

	// if payload is less than head node
	if payload < b.head.payload {
		b.head.left.delete(payload)
		return
	}

	// if payload is greater than head node
	b.head.right.delete(payload)
}

func (n *bstNode[T]) delete(payload T) {

	// if empty node return
	if n == nil {
		return
	}

	// found node to delete
	if n.payload == payload {

		// if both left and right are empty delete node
		if n.left == nil && n.right == nil {
			n = nil
			return
		} else if n.left != nil && n.right == nil {
			// if only left right node is empty
			// move left node up to current node
			n = n.left
			return
		}

		// left and right have nodes
		n = n.left
		return
	}

	// if payload is less than node
	if payload < n.payload {

		if n.left == nil {
			return
		}

		n.left.delete(payload)
		return
	}

	// if payload is greater than node
	if n.right == nil {
		return
	}

	n.right.delete(payload)
}

// FindClosestUint find highest node with uint payload
func (b *BST[T]) FindClosestUint(payload uint) (uint, bool) {

	Tpayload := T(payload)

	n, found := b.findClosest(Tpayload)
	if !found {
		return 0, false
	}

	return uint(n), true
}

// FindClosestUint8 find highest node with uint8 payload
func (b *BST[T]) FindClosestUint8(payload uint8) (uint8, bool) {

	Tpayload := T(payload)

	n, found := b.findClosest(Tpayload)
	if !found {
		return 0, false
	}

	return uint8(n), true
}

// FindClosestUint32 find highest node with uint32 payload
func (b *BST[T]) FindClosestUint32(payload uint32) (uint32, bool) {

	Tpayload := T(payload)

	n, found := b.findClosest(Tpayload)
	if !found {
		return 0, false
	}

	return uint32(n), true
}

// FindClosestUint64 find highest node with uint64 payload
func (b *BST[T]) FindClosestUint64(payload uint64) (uint64, bool) {

	Tpayload := T(payload)

	n, found := b.findClosest(Tpayload)
	if !found {
		return 0, false
	}

	return uint64(n), true
}

// FindClosestInt find highest node with int payload
func (b *BST[T]) FindClosestInt(payload int) (int, bool) {

	Tpayload := T(payload)

	n, found := b.findClosest(Tpayload)
	if !found {
		return 0, false
	}

	return int(n), true
}

// FindClosestInt32 find highest node with int32 payload
func (b *BST[T]) FindClosestInt32(payload int32) (int32, bool) {

	Tpayload := T(payload)

	n, found := b.findClosest(Tpayload)
	if !found {
		return 0, false
	}

	return int32(n), true
}

// FindClosestInt64 find highest node with int64 payload
func (b *BST[T]) FindClosestInt64(payload int64) (int64, bool) {

	Tpayload := T(payload)

	n, found := b.findClosest(Tpayload)
	if !found {
		return 0, false
	}

	return int64(n), true
}

// Between find highest node with any payload
func (b *BST[T]) findClosest(payload T) (T, bool) {

	// if empty tree return false
	if b.head == nil {
		return 0, false
	}

	// if only head node return head node
	if b.head.left == nil && b.head.right == nil {
		return b.head.payload, true
	}

	// if payload is less than head node
	if payload < b.head.payload {

		if b.head.left == nil {
			return b.head.payload, true
		}

		return b.head.left.findClosest(payload)
	}

	// if payload is greater than head node
	if b.head.right == nil {
		return b.head.payload, true
	}

	return b.head.right.findClosest(payload)
}

func (n *bstNode[T]) findClosest(payload T) (T, bool) {

	// if empty node return false
	if n == nil {
		return 0, false
	}

	// if only node return node
	if n.left == nil && n.right == nil {
		return n.payload, true
	}

	// if payload is less than node
	if payload < n.payload {
		if n.left == nil {
			return n.payload, true
		}
		return n.left.findClosest(payload)
	}

	// if payload is greater than node
	if n.right == nil {
		return n.payload, true
	}

	return n.right.findClosest(payload)
}

// Stringify return string representation of tree
func (b *BST[T]) Stringify() []string {
	return b.head.stringify()
}

func (n *bstNode[T]) stringify() []string {

	if n == nil {
		return []string{}
	}

	var nodes []string
	nodes = append(nodes, fmt.Sprintf("%d", n.payload))

	if n.left != nil {
		nodes = append(nodes, n.left.stringify()...)
	}

	if n.right != nil {
		nodes = append(nodes, n.right.stringify()...)
	}

	return nodes
}
