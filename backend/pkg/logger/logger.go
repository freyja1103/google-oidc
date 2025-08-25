package logger

import (
	"context"
	"google-oidc/pkg/logger/sloghandler"
	"google-oidc/pkg/requestid"
)

const skippingFrameCount = 3

func Info(ctx context.Context, msg string, args ...any) {
	requestid.FromContext(sloghandler.WithPC(ctx, skippingFrameCount)).Info(msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	requestid.FromContext(sloghandler.WithPC(ctx, skippingFrameCount)).Error(msg, args...)
}

func Debug(ctx context.Context, msg string, args ...any) {
	requestid.FromContext(sloghandler.WithPC(ctx, skippingFrameCount)).Debug(msg, args...)
}
