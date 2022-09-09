# Format

```log
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
cache L1: 32768
cache L2: 262144
cache L3: 8388608
concurrency: 8
Generating data set with 100000 items
goos: windows
goarch: amd64
pkg: github.com/DaanV2/High-Performance-Cache/benchmarks
cpu: Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz
Benchmark_Duplicate_Items_Write_Test/MapCache_testing:_Writing-8         	  260517	     23746 ns/op	       100.0 N	         0.0003839 N/op	     616 B/op	      11 allocs/op
cache bucket slice size: 1365
panic: runtime error: index out of range [-829]
```