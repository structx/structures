package bench

import (
	"testing"

	"github.com/structx/structs"
)

var (
	queue *structs.Queue[int64] = structs.MustNewQueueInt64[int64]()
)

func BenchmarkEnqueue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		queue.Enqueue(int64(n))
	}
}
