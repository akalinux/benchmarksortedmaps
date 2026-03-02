# Sorted Map Bencharks

## Rankings so far

Write tests:
  1. Go's Native Map
  2. github.com/akalinux/orderedmap
  3. github.com/egregors/sortedmap
  4. github.com/google/btree

Get Single Element Test
  1. Go's Native Map
  2. github.com/egregors/sortedmap
  3. github.com/akalinux/orderedmap
  4. github.com/google/btree

Count Elements between 2 keys
  1. github.com/akalinux/orderedmap
  2. github.com/google/btree
  3. Go's Native Map
  4. github.com/egregors/sortedmap

Mass Key Removal tests
  1. github.com/akalinux/orderedmap
  2. Go's Native Map
  3. github.com/google/btree
  4. github.com/egregors/sortedmap

Full benchmarks
```
go test -bench=. -benchmem -cpu 10
goos: linux
goarch: amd64
pkg: github.com/akalinux/benchmarksortedmaps
cpu: AMD Ryzen 9 9950X 16-Core Processor            
BenchmarkAll/Native_Map_Put:_50-10                483680              2416 ns/op            2593 B/op         54 allocs/op
BenchmarkAll/CenterTree_Put:_50-10                437269              2783 ns/op            3753 B/op         53 allocs/op
BenchmarkAll/sortedmap_Put-10                     198336              5855 ns/op           10773 B/op        118 allocs/op
BenchmarkAll/btree_Put-10                         114285             10156 ns/op           10861 B/op        322 allocs/op
BenchmarkAll/Native_Map_Get:_50-10                594604              1954 ns/op             200 B/op         50 allocs/op
BenchmarkAll/CenterTree_Get:_50-10                367112              3246 ns/op             200 B/op         50 allocs/op
BenchmarkAll/sortedmap_Get:_50-10                 581899              2059 ns/op             200 B/op         50 allocs/op
BenchmarkAll/btree_Get:_50-10                     224839              5219 ns/op            1800 B/op        100 allocs/op
BenchmarkAll/Native_Count:_50-10                 4249303               280.2 ns/op             0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Count:_50-10            76375148                15.79 ns/op            0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Count:_50-10               674720              1788 ns/op            1632 B/op         51 allocs/op
BenchmarkAll/btree_Count:_50-10                  5075013               234.7 ns/op            64 B/op          4 allocs/op
BenchmarkAll/Native_Mass_Remove:_25-10            381058              3156 ns/op             112 B/op         25 allocs/op
BenchmarkAll/CenterTree_Mass_Remove:_25-10       2907139               411.3 ns/op             0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Mass_Remove:_25-10          19539             60646 ns/op           11615 B/op        375 allocs/op
BenchmarkAll/btree_Mass_Remove:_25-10              18643             63916 ns/op            1226 B/op         50 allocs/op
PASS
ok      github.com/akalinux/benchmarksortedmaps 622.629s
```
