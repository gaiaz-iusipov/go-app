package httpclient

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func New(transport http.RoundTripper) *http.Client {
	return &http.Client{
		Transport: RoundTripper{rt: otelhttp.NewTransport(transport,
			otelhttp.WithSpanNameFormatter(spanNameFormatter),
		)},
	}
}
