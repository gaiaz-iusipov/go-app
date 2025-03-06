package httpservermw_test

import (
	"github.com/alecthomas/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"

	httpheader "github.com/gaiaz-iusipov/go-app/http/header"
	httpservermw "github.com/gaiaz-iusipov/go-app/http/server/mw"
)

func TestBasicAuth(t *testing.T) {
	mw := httpservermw.BasicAuth("user", "pass", "realm")
	handler := http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, _ = rw.Write([]byte("ok"))
	})

	mux := http.NewServeMux()
	mux.Handle("/", mw(handler))

	t.Run("without basic auth", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(t.Context(), http.MethodGet, "/", nil)

		mux.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Equal(t, `Basic realm="realm"`, rec.Header().Get(httpheader.WWWAuthenticate))
		assert.Zero(t, rec.Body.String())
	})

	t.Run("with basic auth", func(t *testing.T) {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequestWithContext(t.Context(), http.MethodGet, "/", nil)
		req.SetBasicAuth("user", "pass")

		mux.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Zero(t, rec.Header().Get(httpheader.WWWAuthenticate))
		assert.Equal(t, "ok", rec.Body.String())
	})
}
