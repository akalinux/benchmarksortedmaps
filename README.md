# Sorted Map Bencharks

## Rankings so far

Write tests:
  1. github.com/akalinux/orderedmap
  2. Go's Native Map
  3. github.com/egregors/sortedmap
  4. github.com/google/btree

Get Single Element Test
  1. Go's Native Map
  2. github.com/egregors/sortedmap
  3. github.com/google/btree
  4. github.com/akalinux/orderedmap

Count Elements between 2 keys
  1. github.com/akalinux/orderedmap
  2. github.com/google/btree
  3. Go's Native Map
  4. github.com/egregors/sortedmap

Mass Removal by range
  1. github.com/akalinux/orderedmap
  2. Go's Native Map
  3. github.com/google/btree
  4. github.com/egregors/sortedmap

Full benchmarks
```
BenchmarkAll/Native_Map_Put:_1000
BenchmarkAll/Native_Map_Put:_1000-32            1000000000               0.0000930 ns/op               0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Put:_1000
BenchmarkAll/CenterTree_Put:_1000-32            1000000000               0.0000555 ns/op               0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Put
BenchmarkAll/sortedmap_Put-32                   1000000000               0.0001377 ns/op               0 B/op          0 allocs/op
BenchmarkAll/btree_Put
BenchmarkAll/btree_Put-32                       1000000000               0.0003785 ns/op               0 B/op          0 allocs/op
    benchmarks_test.go:64: ** Read tests
BenchmarkAll/Native_Map_Get:_1000
BenchmarkAll/Native_Map_Get:_1000-32            1000000000               0.0000490 ns/op               0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Get:_1000
BenchmarkAll/CenterTree_Get:_1000-32            1000000000               0.0001047 ns/op               0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Get:_1000
BenchmarkAll/sortedmap_Get:_1000-32             1000000000               0.0000497 ns/op               0 B/op          0 allocs/op
BenchmarkAll/btree_Get:_1000
BenchmarkAll/btree_Get:_1000-32                 1000000000               0.0002362 ns/op               0 B/op          0 allocs/op
    benchmarks_test.go:90: Count all values from mid to end
BenchmarkAll/Native_Count:_1000
BenchmarkAll/Native_Count:_1000-32              1000000000               0.0000345 ns/op               0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Count:_1000
BenchmarkAll/CenterTree_Count:_1000-32          1000000000               0.0000006 ns/op               0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Count:_1000
BenchmarkAll/sortedmap_Count:_1000-32           1000000000               0.0000312 ns/op               0 B/op          0 allocs/op
BenchmarkAll/btree_Count:_1000
BenchmarkAll/btree_Count:_1000-32               1000000000               0.0000463 ns/op               0 B/op          0 allocs/op
    benchmarks_test.go:26: Working with set size: 2000
    benchmarks_test.go:29: ** Write tests
BenchmarkAll/Native_Map_Put:_2000
BenchmarkAll/Native_Map_Put:_2000-32            1000000000               0.0001283 ns/op               0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Put:_2000
BenchmarkAll/CenterTree_Put:_2000-32            1000000000               0.0001326 ns/op               0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Put#01
BenchmarkAll/sortedmap_Put#01-32                1000000000               0.0002168 ns/op               0 B/op          0 allocs/op
BenchmarkAll/btree_Put#01
BenchmarkAll/btree_Put#01-32                    1000000000               0.0006424 ns/op               0 B/op          0 allocs/op
    benchmarks_test.go:64: ** Read tests
BenchmarkAll/Native_Map_Get:_2000
BenchmarkAll/Native_Map_Get:_2000-32            1000000000               0.0001010 ns/op               0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Get:_2000
BenchmarkAll/CenterTree_Get:_2000-32            1000000000               0.0001768 ns/op               0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Get:_2000
BenchmarkAll/sortedmap_Get:_2000-32             1000000000               0.0001363 ns/op               0 B/op          0 allocs/op
BenchmarkAll/btree_Get:_2000
BenchmarkAll/btree_Get:_2000-32                 1000000000               0.0003423 ns/op               0 B/op          0 allocs/op
    benchmarks_test.go:90: Count all values from mid to end
BenchmarkAll/Native_Count:_2000
BenchmarkAll/Native_Count:_2000-32              1000000000               0.0000142 ns/op               0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Count:_2000
BenchmarkAll/CenterTree_Count:_2000-32          1000000000               0.0000007 ns/op               0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Count:_2000
BenchmarkAll/sortedmap_Count:_2000-32           1000000000               0.0000566 ns/op               0 B/op          0 allocs/op
BenchmarkAll/btree_Count:_2000
BenchmarkAll/btree_Count:_2000-32               1000000000               0.0000486 ns/op               0 B/op          0 allocs/op
    benchmarks_test.go:26: Working with set size: 5000
    benchmarks_test.go:29: ** Write tests
BenchmarkAll/Native_Map_Put:_5000
BenchmarkAll/Native_Map_Put:_5000-32            1000000000               0.0002458 ns/op               0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Put:_5000
BenchmarkAll/CenterTree_Put:_5000-32            1000000000               0.0002625 ns/op               0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Put#02
BenchmarkAll/sortedmap_Put#02-32                1000000000               0.0005304 ns/op               0 B/op          0 allocs/op
BenchmarkAll/btree_Put#02
BenchmarkAll/btree_Put#02-32                    1000000000               0.001613 ns/op        0 B/op          0 allocs/op
    benchmarks_test.go:64: ** Read tests
BenchmarkAll/Native_Map_Get:_5000
BenchmarkAll/Native_Map_Get:_5000-32            1000000000               0.0002628 ns/op               0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Get:_5000
BenchmarkAll/CenterTree_Get:_5000-32            1000000000               0.0004620 ns/op               0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Get:_5000
BenchmarkAll/sortedmap_Get:_5000-32             1000000000               0.0003145 ns/op               0 B/op          0 allocs/op
BenchmarkAll/btree_Get:_5000
BenchmarkAll/btree_Get:_5000-32                 1000000000               0.0008712 ns/op               0 B/op          0 allocs/op
    benchmarks_test.go:90: Count all values from mid to end
BenchmarkAll/Native_Count:_5000
BenchmarkAll/Native_Count:_5000-32              1000000000               0.0000315 ns/op               0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Count:_5000
BenchmarkAll/CenterTree_Count:_5000-32          1000000000               0.0000023 ns/op               0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Count:_5000
BenchmarkAll/sortedmap_Count:_5000-32           1000000000               0.0001869 ns/op               0 B/op          0 allocs/op
BenchmarkAll/btree_Count:_5000
BenchmarkAll/btree_Count:_5000-32               1000000000               0.0000230 ns/op               0 B/op          0 allocs/op
    benchmarks_test.go:124: Delete first 999 elements
BenchmarkAll/Native_Mass_Remove:_5000
BenchmarkAll/Native_Mass_Remove:_5000-32        1000000000               0.0000964 ns/op               0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Mass_Remove:_5000
BenchmarkAll/CenterTree_Mass_Remove:_5000-32    1000000000               0.0000006 ns/op               0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Mass_Remove:_5000
BenchmarkAll/sortedmap_Mass_Remove:_5000-32     1000000000               0.07483 ns/op         0 B/op          0 allocs/op
BenchmarkAll/btree_Count:_5000#01-32            1000000000               0.0001720 ns/op               0 B/op          0 allocs/op
```
