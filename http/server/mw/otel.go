package httpservermw

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
)

func RoutePattern(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		otelhttp.WithRouteTag(req.Pattern, next).ServeHTTP(rw, req)
	})
}

func TraceIDHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if traceID := trace.SpanContextFromContext(req.Context()).TraceID(); traceID.IsValid() {
			rw.Header().Add("X-Trace-Id", traceID.String())
		}
		next.ServeHTTP(rw, req)
	})
}
