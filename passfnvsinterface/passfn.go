package passfnvsinterface

import "context"

func PassFn(ctx context.Context, fn func(context.Context, string) error) error {
	return fn(ctx, "test value")
}
