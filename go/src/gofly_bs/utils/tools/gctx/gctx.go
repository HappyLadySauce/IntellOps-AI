package gctx

import (
	"context"

	"gofly/utils/tools/gtrace"
)

type (
	Ctx    = context.Context // Ctx is short name alias for context.Context.
	StrKey string            // StrKey is a type for warps basic type string as context key.
)

// New creates and returns a context which contains context id.
func New() context.Context {
	return WithCtx(context.Background())
}

// WithCtx creates and returns a context containing context id upon given parent context `ctx`.
func WithCtx(ctx context.Context) context.Context {
	if CtxId(ctx) != "" {
		return ctx
	}
	var span *gtrace.Span
	ctx, span = gtrace.NewSpan(ctx, "gctx.WithCtx")
	defer span.End()
	return ctx
}

// CtxId retrieves and returns the context id from context.
func CtxId(ctx context.Context) string {
	return gtrace.GetTraceID(ctx)
}
