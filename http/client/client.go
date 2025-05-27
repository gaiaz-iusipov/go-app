package httpclient

import (
	"context"
	"net/http"
	"net/http/httptrace"

	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// New wraps the provided [http.RoundTripper] and returns new [http.Client].
//
// If the provided [http.RoundTripper] is nil, [http.DefaultTransport] will be used
// as the base [http.RoundTripper].
func New(transport http.RoundTripper) *http.Client {
	return &http.Client{
		Transport: RoundTripper{rt: otelhttp.NewTransport(transport,
			otelhttp.WithSpanNameFormatter(spanNameFormatter),
			otelhttp.WithClientTrace(func(ctx context.Context) *httptrace.ClientTrace {
				return otelhttptrace.NewClientTrace(ctx)
			}),
		)},
	}
}
