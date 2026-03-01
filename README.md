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

Mass Remove the first 999 elements
  1. github.com/akalinux/orderedmap
  2. Go's Native Map
  3. github.com/google/btree
  4. github.com/egregors/sortedmap

Full benchmarks
```
go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/akalinux/benchmarksortedmaps
cpu: AMD Ryzen 9 9950X 16-Core Processor            
BenchmarkAll/Native_Map_Put:_50-32                435003              2560 ns/op            2596 B/op         54 allocs/op
BenchmarkAll/CenterTree_Put:_50-32                395616              2974 ns/op            3758 B/op         53 allocs/op
BenchmarkAll/sortedmap_Put-32                     181574              6503 ns/op           10787 B/op        118 allocs/op
BenchmarkAll/btree_Put-32                         111808             10710 ns/op           10874 B/op        322 allocs/op
BenchmarkAll/Native_Map_Get:_50-32                589975              1986 ns/op             200 B/op         50 allocs/op
BenchmarkAll/CenterTree_Get:_50-32                358390              3340 ns/op             200 B/op         50 allocs/op
BenchmarkAll/sortedmap_Get:_50-32                 611535              1962 ns/op             200 B/op         50 allocs/op
BenchmarkAll/btree_Get:_50-32                     220504              5438 ns/op            1802 B/op        100 allocs/op
BenchmarkAll/Native_Count:_50-32                 4377718               272.6 ns/op             0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Count:_50-32            66874784                17.87 ns/op            0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Count:_50-32               628212              1881 ns/op            1632 B/op         51 allocs/op
BenchmarkAll/btree_Count:_50-32                  5046243               237.1 ns/op            64 B/op          4 allocs/op
BenchmarkAll/Native_Mass_Remove:_25-32            393682              3134 ns/op             112 B/op         25 allocs/op
BenchmarkAll/CenterTree_Mass_Remove:_25-32       3256267               379.9 ns/op             0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Mass_Remove:_25-32          19072             60802 ns/op           11626 B/op        375 allocs/op
BenchmarkAll/btree_Mass_Remove:_25-32              18164             63968 ns/op            1219 B/op         50 allocs/op
PASS
ok      github.com/akalinux/benchmarksortedmaps 660.500s
```
