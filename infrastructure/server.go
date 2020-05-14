package infrastructure

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nekochans/kimono-app-api/infrastructure/httputil"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type HTTPServer struct {
	logger *zap.Logger
	router *chi.Mux
}

func NewServer(logger *zap.Logger) *HTTPServer {
	return &HTTPServer{
		router: chi.NewRouter(),
		logger: logger,
	}
}

func (s *HTTPServer) Middleware() {
	s.router.Use(httputil.Log(s.logger))
}

func (s *HTTPServer) Router() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, HTTPサーバ")
	})
}

func StartHTTPServer() {
	logOpts := LoggerOptions{}
	logger, err := NewLoggerFromOptions(logOpts)
	defer logger.Sync()

	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to create logger: %s\n", err)
		os.Exit(1)
	}
	s := NewServer(logger)
	s.Middleware()
	s.Router()
	log.Println("Starting app")
	_ = http.ListenAndServe(":8888", s.router)
}
