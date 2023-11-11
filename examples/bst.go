// Package examples showcase basic usage of structure or algorithm
package examples

import (
	"fmt"

	"github.com/strcutx/structs"
)

// ExampleBst example binary search tree
func ExampleBst() {

	// create new binary search tree
	tree := structs.NewBST[int64]()

	// add some nodes to the tree
	tree.AddNodeInt64(3)
	tree.AddNodeInt64(74)
	tree.AddNodeInt64(2)

	// search for highest node closest
	// to argument : 1
	p, m := tree.FindClosestInt64(1)
	if !m {
		// failed to find match
		fmt.Println("failed to find match")
	}

	// p : matched payload
	// expected output: 2
	fmt.Printf("matched payload %d", p)
}
