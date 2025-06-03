package httpservermw

import (
	"net/http"

	httpheader "github.com/gaiaz-iusipov/go-app/http/header"
)

func Header(key, val string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add(key, val)
			next.ServeHTTP(rw, req)
		})
	}
}

func HeaderImmutable(next http.Handler) http.Handler {
	return Header(httpheader.CacheControl, httpheader.CacheControlImmutable)(next)
}
