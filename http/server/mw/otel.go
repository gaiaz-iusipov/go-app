package httpservermw

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
)

func ChiRouteTag(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		routePattern := chi.RouteContext(req.Context()).RoutePattern()
		if routePattern == "" {
			routePattern = "_other"
		}
		otelhttp.WithRouteTag(routePattern, next).ServeHTTP(rw, req)
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
