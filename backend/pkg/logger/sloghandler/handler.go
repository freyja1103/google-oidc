package sloghandler

import (
	"context"
	"io"
	"log/slog"
)

var _ slog.Handler = (*Handler)(nil)

type Handler struct {
	handler slog.Handler
}

func NewHandler(w io.Writer, opts *slog.HandlerOptions) *Handler {
	return &Handler{
		handler: slog.NewJSONHandler(w, opts),
	}
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	pc, ok := ctx.Value(counterKey).(uintptr)
	if ok {
		r.PC = pc
	}
	return h.handler.Handle(ctx, r)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{
		handler: h.handler.WithAttrs(attrs),
	}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		handler: h.handler.WithGroup(name),
	}
}
