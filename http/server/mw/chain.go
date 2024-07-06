package httpservermw

import (
	"net/http"
)

type Chain []func(next http.Handler) http.Handler

func (c Chain) Middleware(handler http.Handler) http.Handler {
	for i := len(c) - 1; i >= 0; i-- { // iterate in reverse order
		handler = c[i](handler)
	}
	return handler
}
