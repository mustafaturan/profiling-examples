# Slice of value vs pointer

Are you using slice of value or pointer? Checkout [benchmark_test.go](./benchmark_test.go) for implementation details. Here is the results:

```
go test -run none -bench=. -benchtime 10000000x
goos: darwin
goarch: amd64
pkg: github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer
cpu: Intel(R) Core(TM) i5-6267U CPU @ 2.90GHz
BenchmarkSliceOfValue-4     	10000000	         6.549 ns/op	       0 B/op	       0 allocs/op
BenchmarkSliceOfPointer-4   	10000000	       994.0 ns/op	    4265 B/op	       0 allocs/op
PASS
ok  	github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer	10.185s
```

Memory profiling results
```
Showing nodes accounting for 39.80GB, 100% of 39.81GB total
Dropped 17 nodes (cum <= 0.20GB)
      flat  flat%   sum%        cum   cum%
   39.80GB   100%   100%    39.80GB   100%  github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer.SliceOfPointer (inline)
         0     0%   100%    39.80GB   100%  github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer_test.BenchmarkSliceOfPointer
         0     0%   100%    39.80GB   100%  testing.(*B).launch
         0     0%   100%    39.80GB   100%  testing.(*B).runN
```

## Notes from Go language official doc

[When are function parameters passed by value?](https://golang.org/doc/faq#pass_by_value)

> Map and slice values behave like pointers: they are descriptors that contain pointers to the underlying map or slice data. Copying a map or slice value doesn't copy the data it points to.

[How do I know whether a variable is allocated on the heap or the stack?](https://golang.org/doc/faq#stack_or_heap)

> If a local variable is very large, it might make more sense to store it on the heap rather than the stack.
