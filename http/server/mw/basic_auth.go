package httpservermw

import (
	"fmt"
	"net/http"
)

// BasicAuth implements a simple middleware for adding HTTP Basic Authentication.
func BasicAuth(username, password, realm string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			reqUsername, reqPassword, ok := req.BasicAuth()
			if !ok || username != reqUsername || password != reqPassword {
				rw.Header().Add("WWW-Authenticate", fmt.Sprintf("Basic realm=%q", realm))
				rw.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(rw, req)
		})
	}
}
