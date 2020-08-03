package httputil

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

func Log(l *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				l.Info("Access",
					zap.String("request_id", middleware.GetReqID(r.Context())),
					zap.String("request_remote", r.RemoteAddr),
					zap.String("request_protocol", r.Proto),
					zap.String("request_method", r.Method),
					zap.String("request_url", r.RequestURI),
					zap.String("request_path", r.URL.Path),
					zap.String("request_host", r.Host),
					zap.String("request_referer", r.Referer()),
					zap.Int("responses_status", ww.Status()),
					zap.Int("response_length", ww.BytesWritten()),
					zap.Duration("duration", time.Since(t1)),
				)
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
