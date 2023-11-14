## Binary Seach Tree (BST)

```bash
goos: linux
goarch: amd64
pkg: github.com/structx/structs/bench
cpu: Intel(R) Core(TM) i5-10300H CPU @ 2.50GHz
=== RUN   BenchmarkBSTAddNode
BenchmarkBSTAddNode
BenchmarkBSTAddNode-8              18493            112230 ns/op              23 B/op          1 allocs/op
PASS
ok      github.com/structx/structs/bench        2.731s
```

## Queue

```bash
goos: linux
goarch: amd64
pkg: github.com/structx/structs/bench
cpu: Intel(R) Core(TM) i5-10300H CPU @ 2.50GHz
=== RUN   BenchmarkEnqueue
BenchmarkEnqueue
BenchmarkEnqueue-8       2429272               553.8 ns/op           157 B/op          1 allocs/op
PASS
ok      github.com/structx/structs/bench        2.362s
```

## Consistent Hashing (Weightless)

```bash
goos: linux
goarch: amd64
pkg: github.com/structx/structs/bench
cpu: Intel(R) Core(TM) i5-10300H CPU @ 2.50GHz
=== RUN   BenchmarkRingAddPoint
BenchmarkRingAddPoint
BenchmarkRingAddPoint-8          1000000            158411 ns/op              72 B/op          4 allocs/op
PASS
ok      github.com/structx/structs/bench        159.024s
```