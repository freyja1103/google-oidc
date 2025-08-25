package requestid

import (
	"context"
	"log/slog"
)

type loggerKey struct{}

func newContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	log, ok := ctx.Value(loggerKey{}).(*slog.Logger)
	if !ok {
		return nil
	}
	return log
}
