# Structures 

package is under active development. `LinkedList` and `DoubleLinkedList` are the next data structures to be added.

Generic golang data structures and algorithms to be easily integrated into more complex systems. Data structures are designed to be generic and easily integrated with algorithms within this package. This package is designed to be used as building blocks to create a distributed network. 

For example the `ConsistentHashRingWeightless` uses the `BST` for a quick lookup of nodes. 

## Usage

```bash
go get github.com/structx/structures
```

## Docs

This package has a wiki page to provide documentation on each data structure and algorithm. This repository is a learning opportunity for me
so I will do my best to explain the structs and algs. 

## Examples

please view source code under `/examples` to see how to implement a binary search tree. 

```golang

    // create new tree
    tree := structs.NewBST()

    // add node to tree
    tree.AddNodeInt64(1) // will add a single node to the tree

    // you can remove and find closest node
    tree.FindClosestInt64(2)
    tree.DeleteInt64(1)
```

## Benchmarks

Each struct and alg will have a benchmark test listed under the `/bench` dir. I run the benchmark tests on my laptop so please feel free to run on your own machines and share results with me. 