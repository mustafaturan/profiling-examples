# Passing function vs passing interface as arg

Functions are first class citizens in Go language. Is it good idea to use pure functions instead of interface types?

Checkout [benchmark_test.go](./benchmark_test.go) for implementation details. Here is the results:

```
go test -run none -bench=. -benchtime 10000000x
goos: darwin
goarch: amd64
pkg: github.com/mustafaturan/profiling-examples/passfnvsinterface
cpu: Intel(R) Core(TM) i5-6267U CPU @ 2.90GHz
BenchmarkPassFn-4                    	10000000	         0.5964 ns/op
BenchmarkPassLocalFn-4               	10000000	         0.5762 ns/op
BenchmarkPassStructFn-4              	10000000	         5.019 ns/op
BenchmarkPassStructFnVar-4           	10000000	         4.004 ns/op
BenchmarkPassIntefaceFn-4            	10000000	         5.006 ns/op
BenchmarkPassIntefaceFnVar-4         	10000000	         4.224 ns/op
BenchmarkPassInterfaceStructDoer-4   	10000000	         4.052 ns/op
BenchmarkPassInterfaceDoer-4         	10000000	         4.749 ns/op
PASS
ok  	github.com/mustafaturan/profiling-examples/passfnvsinterface	0.503s
```
