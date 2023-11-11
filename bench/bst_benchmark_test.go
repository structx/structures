package bench

import (
	"testing"

	"github.com/strcutx/structs"
)

var (
	tree *structs.BST[int] = structs.NewBST[int]()
)

func init() {
	for i := 0; i < 10000; i++ {
		tree.AddNodeInt(i)
	}
}

func BenchmarkBSTAddNode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tree.AddNodeInt(n)
	}
}

func BenchmarkBSTFindClosest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tree.FindClosestInt(n)
	}
}

func BenchmarkBSTDelete(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tree.DeleteInt(n)
	}
}
