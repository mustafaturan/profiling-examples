package towardstozeroallocation

import (
	"testing"

	"github.com/mustafaturan/profiling-examples/towardstozeroallocation/version1"
	"github.com/mustafaturan/profiling-examples/towardstozeroallocation/version2"
	"github.com/mustafaturan/profiling-examples/towardstozeroallocation/version3"
)

func BenchmarkV1Next(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = version1.Next()
	}
}

func BenchmarkV2Next(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = version2.Next()
	}
}

func BenchmarkV3Next(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = version3.Next()
	}
}
