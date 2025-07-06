package httpservermw

import (
	"net/http"
	"path/filepath"

	"go.opentelemetry.io/otel/trace"

	httpheader "github.com/gaiaz-iusipov/go-app/http/header"
)

type Header struct{}

func (Header) Add(key, val string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add(key, val)
			next.ServeHTTP(rw, req)
		})
	}
}

func (h Header) Immutable(next http.Handler) http.Handler {
	return h.Add(httpheader.CacheControl, httpheader.CacheControlImmutable)(next)
}

func (Header) ContentTypeByExt(contentTypeByExt map[string]string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ext := filepath.Ext(req.URL.Path)
			if v, ok := contentTypeByExt[ext]; ok {
				rw.Header().Set(httpheader.ContentType, v)
			}

			next.ServeHTTP(rw, req)
		})
	}
}

func (Header) TraceID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if traceID := trace.SpanContextFromContext(req.Context()).TraceID(); traceID.IsValid() {
			rw.Header().Add(httpheader.TraceID, traceID.String())
		}
		next.ServeHTTP(rw, req)
	})
}
