package httpservermw

import (
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

func TraceIDHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if traceID := trace.SpanContextFromContext(req.Context()).TraceID(); traceID.IsValid() {
			rw.Header().Add("X-Trace-Id", traceID.String())
		}
		next.ServeHTTP(rw, req)
	})
}
