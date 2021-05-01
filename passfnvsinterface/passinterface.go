package passfnvsinterface

import "context"

type Doer interface {
	Do(ctx context.Context, val string) error
}

func PassInterface(ctx context.Context, doer Doer) error {
	return doer.Do(ctx, "test value")
}
