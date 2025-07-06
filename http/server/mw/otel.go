package httpservermw

import (
	"net/http"
	"strings"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

type OTEL struct{}

func (OTEL) RouteMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if idx := strings.IndexByte(req.Pattern, '/'); idx >= 0 {
			labeler, _ := otelhttp.LabelerFromContext(req.Context())
			labeler.Add(semconv.HTTPRoute(req.Pattern[idx:]))
		}
		next.ServeHTTP(rw, req)
	})
}
