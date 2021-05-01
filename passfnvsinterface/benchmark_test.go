package passfnvsinterface_test

import (
	"context"
	"testing"

	"github.com/mustafaturan/profiling-examples/passfnvsinterface"
)

type doer struct{}
type idoer struct{}

func (d doer) Do(ctx context.Context, s string) error {
	return nil
}

func (d *idoer) Do(ctx context.Context, s string) error {
	return nil
}

func do(ctx context.Context, s string) error {
	return nil
}

func BenchmarkPassFn(b *testing.B) {
	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		_ = passfnvsinterface.PassFn(ctx, do)
	}
}

func BenchmarkPassLocalFn(b *testing.B) {
	ctx := context.Background()

	d := func(ctx context.Context, s string) error {
		return nil
	}

	for i := 0; i < b.N; i++ {
		_ = passfnvsinterface.PassFn(ctx, d)
	}
}

func BenchmarkPassStructFn(b *testing.B) {
	ctx := context.Background()
	d := doer{}

	for i := 0; i < b.N; i++ {
		_ = passfnvsinterface.PassFn(ctx, d.Do)
	}
}

func BenchmarkPassStructFnVar(b *testing.B) {
	ctx := context.Background()
	d := doer{}.Do

	for i := 0; i < b.N; i++ {
		_ = passfnvsinterface.PassFn(ctx, d)
	}
}

func BenchmarkPassIntefaceFn(b *testing.B) {
	ctx := context.Background()
	d := &idoer{}

	for i := 0; i < b.N; i++ {
		_ = passfnvsinterface.PassFn(ctx, d.Do)
	}
}

func BenchmarkPassIntefaceFnVar(b *testing.B) {
	ctx := context.Background()
	d := &idoer{}
	do := d.Do

	for i := 0; i < b.N; i++ {
		_ = passfnvsinterface.PassFn(ctx, do)
	}
}

func BenchmarkPassInterfaceStructDoer(b *testing.B) {
	ctx := context.Background()
	d := doer{}

	for i := 0; i < b.N; i++ {
		_ = passfnvsinterface.PassInterface(ctx, d)
	}
}

func BenchmarkPassInterfaceDoer(b *testing.B) {
	ctx := context.Background()
	d := &idoer{}

	for i := 0; i < b.N; i++ {
		_ = passfnvsinterface.PassInterface(ctx, d)
	}
}
