package utils

import (
	"context"
	"log/slog"
)

type NoopHandler struct {
}

func (h *NoopHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (h *NoopHandler) Handle(ctx context.Context, lr slog.Record) error {
	return nil
}

func (h *NoopHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *NoopHandler) WithGroup(name string) slog.Handler {
	return h
}

func (h *NoopHandler) Reset() {
}

func NewNoopHandler() *NoopHandler {
	return &NoopHandler{}
}

var _ slog.Handler = (*NoopHandler)(nil)
