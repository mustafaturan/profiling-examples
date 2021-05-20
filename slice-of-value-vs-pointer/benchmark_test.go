package sliceofvaluevspointer_test

import (
	"testing"

	sliceof "github.com/mustafaturan/profiling-examples/slice-of-value-vs-pointer"
)

func BenchmarkSliceOfValue(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_ = sliceof.SliceOfValue(n % 1_000)
	}
}

func BenchmarkSliceOfValueWithInt(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_ = sliceof.SliceOfValueWithInt(n % 1_000)
	}
}

func BenchmarkSliceOfValueWithLargeData(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_ = sliceof.SliceOfValueWithLargeData(n % 1_000)
	}
}

func BenchmarkSliceOfPointer(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_ = sliceof.SliceOfPointer(n % 1_000)
	}
}

func BenchmarkSliceOfPointerWithInt(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_ = sliceof.SliceOfPointerWithInt(n % 1_000)
	}
}

func BenchmarkSliceOfPointerWithLargeData(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		_ = sliceof.SliceOfPointerWithLargeData(n % 1_000)
	}
}
