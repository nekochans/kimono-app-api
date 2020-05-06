package httputil

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/nekochans/kimono-app-api/infrastructure"
)

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logOpts := infrastructure.LoggerOptions{}
		logger, err := infrastructure.NewLoggerFromOptions(logOpts)
		defer logger.Sync()

		if err != nil {
			fmt.Fprintf(os.Stderr, "[ERROR] Failed to create logger: %s\n", err)
			os.Exit(1)
		}

		t1 := time.Now()
		defer func() {
			logger.Info("Access",
				zap.String("request_id", "-"), // TODO RequestIDを追加する
				zap.String("request_remote", r.RemoteAddr),
				zap.String("request_protocol", r.Proto),
				zap.String("request_method", r.Method),
				zap.String("request_url", r.RequestURI),
				zap.String("request_path", r.URL.Path),
				zap.String("request_host", r.Host),
				zap.String("request_referer", r.Referer()),
				zap.Duration("duration", time.Since(t1)),
				// TODO レスポンスについてもログを出力する
			)
		}()

		h.ServeHTTP(w, r)
	})
}
