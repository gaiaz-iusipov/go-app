package httpclient

import (
	"context"
	"net/http/httptrace"

	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type config struct {
	otelOpts []otelhttp.Option
}

var defaultConfig = config{
	otelOpts: []otelhttp.Option{
		otelhttp.WithSpanNameFormatter(spanNameFormatter),
		otelhttp.WithClientTrace(func(ctx context.Context) *httptrace.ClientTrace {
			return otelhttptrace.NewClientTrace(ctx)
		}),
	},
}

type Option func(cfg *config)

func WithOTELOptions(otelOpts ...otelhttp.Option) Option {
	return func(cfg *config) {
		cfg.otelOpts = append(cfg.otelOpts, otelOpts...)
	}
}
