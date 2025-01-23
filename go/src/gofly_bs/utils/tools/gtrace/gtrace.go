package gtrace

import (
	"context"
	"gofly/utils/tools/gipv4"
	"os"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	tracingCommonKeyIpIntranet        = `ip.intranet`
	tracingCommonKeyIpHostname        = `hostname`
	commandEnvKeyForMaxContentLogSize = "gf.gtrace.max.content.log.size" // To avoid too big tracing content.
	commandEnvKeyForTracingInternal   = "gf.gtrace.tracing.internal"     // For detailed controlling for tracing content.
)

type TracerProvider struct {
	*sdkTrace.TracerProvider
}

var (
	intranetIps, _  = gipv4.GetIntranetIpArray()
	intranetIpStr   = strings.Join(intranetIps, ",")
	hostname, _     = os.Hostname()
	tracingInternal = true // tracingInternal enables tracing for internal type spans.
)

// IsUsingDefaultProvider checks and return if currently using default trace provider.
func IsUsingDefaultProvider() bool {
	_, ok := otel.GetTracerProvider().(*TracerProvider)
	return ok
}

// IsTracingInternal returns whether tracing spans of internal components.
func IsTracingInternal() bool {
	return tracingInternal
}

// CommonLabels returns common used attribute labels:
// ip.intranet, hostname.
func CommonLabels() []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String(tracingCommonKeyIpHostname, hostname),
		attribute.String(tracingCommonKeyIpIntranet, intranetIpStr),
		semconv.HostNameKey.String(hostname),
	}
}

// GetTraceID retrieves and returns TraceId from context.
// It returns an empty string is tracing feature is not activated.
func GetTraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	traceID := trace.SpanContextFromContext(ctx).TraceID()
	if traceID.IsValid() {
		return traceID.String()
	}
	return ""
}

// Span warps trace.Span for compatibility and extension.
type Span struct {
	trace.Span
}

// NewSpan creates a span using default tracer.
func NewSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, *Span) {
	ctx, span := NewTracer().Start(ctx, spanName, opts...)
	return ctx, &Span{
		Span: span,
	}
}
