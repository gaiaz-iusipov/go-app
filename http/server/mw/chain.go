package httpservermw

import (
	"net/http"
	"slices"
)

type Chain []func(next http.Handler) http.Handler

func (c Chain) Middleware(handler http.Handler) http.Handler {
	for _, v := range slices.Backward(c) {
		handler = v(handler)
	}
	return handler
}
