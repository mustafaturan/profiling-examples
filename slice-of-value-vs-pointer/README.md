# Slice of value vs pointer

Are you using slice of value or pointer? Checkout [benchmark_test.go](./benchmark_test.go) for implementation details. Here is the results:

```
goos: darwin
goarch: amd64
pkg: github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer
cpu: Intel(R) Core(TM) i5-6267U CPU @ 2.90GHz
BenchmarkSliceOfValue-4                         10000000                 6.247 ns/op               0 B/op          0 allocs/op
BenchmarkSliceOfValueWithInt-4                  10000000               704.9 ns/op      4265 B/op          0 allocs/op
BenchmarkSliceOfValueWithLargeData-4            10000000              5418 ns/op       25950 B/op          0 allocs/op
BenchmarkSliceOfPointer-4                       10000000              1110 ns/op        4265 B/op          0 allocs/op
BenchmarkSliceOfPointerWithInt-4                10000000              1103 ns/op        4265 B/op          0 allocs/op
BenchmarkSliceOfPointerWithLargeData-4          10000000              1139 ns/op        4265 B/op          0 allocs/op
PASS
ok      github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer94.987s
```

Memory profiling results
```
Showing top 10 nodes out of 12
      flat  flat%   sum%        cum   cum%
  241.34GB 60.30% 60.30%   241.34GB 60.30%  github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer.SliceOfValueWithLargeData (inline)
   39.88GB  9.96% 70.27%    39.88GB  9.96%  github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer.SliceOfPointer (inline)
   39.76GB  9.93% 80.20%    39.76GB  9.93%  github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer.SliceOfPointerWithInt (inline)
   39.62GB  9.90% 90.10%    39.62GB  9.90%  github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer.SliceOfPointerWithLargeData (inline)
   39.61GB  9.90%   100%    39.61GB  9.90%  github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer.SliceOfValueWithInt (inline)
         0     0%   100%    39.88GB  9.96%  github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer_test.BenchmarkSliceOfPointer
         0     0%   100%    39.76GB  9.93%  github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer_test.BenchmarkSliceOfPointerWithInt
         0     0%   100%    39.62GB  9.90%  github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer_test.BenchmarkSliceOfPointerWithLargeData
         0     0%   100%    39.61GB  9.90%  github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer_test.BenchmarkSliceOfValueWithInt
         0     0%   100%   241.34GB 60.30%  github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer_test.BenchmarkSliceOfValueWithLargeData
```

## TODO

GC pause profiling

## Notes from Go language official doc

[When are function parameters passed by value?](https://golang.org/doc/faq#pass_by_value)

> Map and slice values behave like pointers: they are descriptors that contain pointers to the underlying map or slice data. Copying a map or slice value doesn't copy the data it points to.

[How do I know whether a variable is allocated on the heap or the stack?](https://golang.org/doc/faq#stack_or_heap)

> If a local variable is very large, it might make more sense to store it on the heap rather than the stack.
