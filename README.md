# Sorted Map Bencharks

Full benchmarks
```
go test -bench=. -benchmem -v -cpu 10
goos: linux
goarch: amd64
pkg: github.com/akalinux/benchmarksortedmaps
cpu: AMD Ryzen 9 9950X 16-Core Processor            
BenchmarkAll
    benchmarks_test.go:26: Working with set size: 1000
    benchmarks_test.go:29: ** Write tests
BenchmarkAll/Native_Map_Put
BenchmarkAll/Native_Map_Put-10             22053             54295 ns/op           95023 B/op       1750 allocs/op
BenchmarkAll/CenterTree_Put
BenchmarkAll/CenterTree_Put-10             21289             56550 ns/op           78598 B/op       1747 allocs/op
BenchmarkAll/sortedmap_Put
BenchmarkAll/sortedmap_Put-10               8084            137616 ns/op          275718 B/op       2779 allocs/op
BenchmarkAll/btree_Put
BenchmarkAll/btree_Put-10                   4402            275982 ns/op          242970 B/op       7688 allocs/op
    benchmarks_test.go:86: ** Single Read tests
BenchmarkAll/Native_Map_Ge
BenchmarkAll/Native_Map_Ge-10              27411             44073 ns/op           12934 B/op       1744 allocs/op
BenchmarkAll/CenterTree_Get
BenchmarkAll/CenterTree_Get-10             14163             83630 ns/op           12935 B/op       1744 allocs/op
BenchmarkAll/sortedmap_Get
BenchmarkAll/sortedmap_Get-10              26869             44370 ns/op           12935 B/op       1744 allocs/op
BenchmarkAll/btree_Get
BenchmarkAll/btree_Get-10                   8070            146664 ns/op           44953 B/op       2744 allocs/op
    benchmarks_test.go:121: Count all values from begin to mid
BenchmarkAll/Native_Count
BenchmarkAll/Native_Count-10              150838              7693 ns/op               0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Count
BenchmarkAll/CenterTree_Count-10        44672611                26.93 ns/op            0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Count
BenchmarkAll/sortedmap_Count-10         28128948                42.79 ns/op           64 B/op          2 allocs/op
BenchmarkAll/btree_Count
BenchmarkAll/btree_Count-10             28563042                42.21 ns/op           40 B/op          2 allocs/op
    benchmarks_test.go:165: Find and delete first half of all elements
BenchmarkAll/Native_Mass_Remove
BenchmarkAll/Native_Mass_Remove-10       2091104               571.4 ns/op             0 B/op          0 allocs/op
BenchmarkAll/CenterTree_Mass_Remove
BenchmarkAll/CenterTree_Mass_Remove-10   2803908               429.6 ns/op             0 B/op          0 allocs/op
BenchmarkAll/sortedmap_Mass_Remove
BenchmarkAll/sortedmap_Mass_Remove-10        328           3693223 ns/op         8085054 B/op     252771 allocs/op
BenchmarkAll/btree_Mass_Remove
BenchmarkAll/btree_Mass_Remove-10         494234              2423 ns/op              32 B/op          1 allocs/op
PASS
ok      github.com/akalinux/benchmarksortedmaps 1231.189s
```
