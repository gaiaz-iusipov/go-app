package httpclient

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
)

func spanNameFormatter(operation string, req *http.Request) string {
	if requestName := requestNameFromContext(req.Context()); requestName != "" {
		return requestName
	}
	return operation
}

type RoundTripper struct {
	rt http.RoundTripper
}

func (rt RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	if requestName := requestNameFromContext(ctx); requestName != "" {
		labeler, _ := otelhttp.LabelerFromContext(ctx)
		labeler.Add(attribute.String("http.request_name", requestName))
		ctx = otelhttp.ContextWithLabeler(ctx, labeler)
		req = req.WithContext(ctx)
	}

	return rt.rt.RoundTrip(req)
}
