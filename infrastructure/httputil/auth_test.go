package httputil

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		name               string
		request            func() *http.Request
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			"OK response when including the Authorization header",
			func() *http.Request {
				req, _ := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/", nil)
				req.Header.Add("Authorization", "aaa.aaa.aaa")
				req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
				return req
			},
			"Hello, HTTPサーバ",
			http.StatusOK,
		},
		{
			"Error response when the Authorization header is not included",
			func() *http.Request {
				req, _ := http.NewRequestWithContext(context.TODO(), http.MethodGet, "/", nil)
				req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64)")
				return req
			},
			"Unauthorized",
			http.StatusUnauthorized,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := chi.NewRouter()

			// TODO テスト用のパラメータを渡す
			r.Use(Auth("", "", ""))
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Hello, HTTPサーバ")
			})

			r.ServeHTTP(w, test.request())

			if w.Body.String() != test.expectedResponse {
				t.Error("response Body was not the expected value")
			}
			if w.Code != test.expectedStatusCode {
				t.Error("status code was not the expected value")
			}
		})
	}
}
