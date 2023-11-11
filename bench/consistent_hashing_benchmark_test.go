package bench

import (
	"fmt"
	"testing"

	"github.com/structx/structs/algs"
)

var (
	ringWeightless *algs.ConsistentHashRingWeightless[uint64] = algs.NewConsistentHashRingWeightless[uint64]()
)

func BenchmarkRingAddPoint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ringWeightless.AddPoint(fmt.Sprintf("10.1.%d.1", n))
	}
}
