package sloghandler

import (
	"context"
	"runtime"
)

type counterKeyType struct{}

var counterKey = counterKeyType{}

func WithPC(ctx context.Context, skipCount int) context.Context {
	var pcs [1]uintptr
	runtime.Callers(skipCount, pcs[:])
	return context.WithValue(ctx, counterKey, pcs[0])
}
