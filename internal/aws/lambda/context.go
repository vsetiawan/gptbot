package lambda

import (
	"context"
	"time"
)

const (
	contextTimeout = 5 * time.Second
)

func newDefaultContext() (context.Context, context.CancelFunc) {
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, contextTimeout)
	return ctx, cancelFunc
}
