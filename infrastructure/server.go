package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nekochans/kimono-app-api/infrastructure/httputil"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	logger *zap.Logger
	router *chi.Mux
}

type UserAddress struct {
	Sub  string
	Name string
}

var authenticatedUser string

func NewServer(logger *zap.Logger) *HTTPServer {
	return &HTTPServer{
		router: chi.NewRouter(),
		logger: logger,
	}
}

func (s *HTTPServer) Middleware() {
	s.router.Use(middleware.RequestID)
	s.router.Use(httputil.Log(s.logger))
}

func (s *HTTPServer) Router() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, HTTPサーバ")
	})
	s.router.Get(
		"/account/v1/user/"+authenticatedUser+"/addresses",
		func(w http.ResponseWriter, r *http.Request) {
			userAddress := UserAddress{
				Sub:  authenticatedUser,
				Name: "authenticated user",
			}
			response, _ := json.Marshal(userAddress)
			fmt.Fprint(w, string(response))
		})
	s.router.Get(
		"/account/v1/user/123456789012345678901234567890abcdef/addresses",
		func(w http.ResponseWriter, r *http.Request) {
			userAddress := UserAddress{
				Sub:  "123456789012345678901234567890abcdef",
				Name: "unauthenticated user",
			}
			response, _ := json.Marshal(userAddress)
			fmt.Fprint(w, string(response))
		})
}

func StartHTTPServer() {
	logOpts := LoggerOptions{}
	logger, err := NewLoggerFromOptions(logOpts)
	authenticatedUser = os.Getenv("AUTHENTICATED_USER")

	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Fatal(err, "logger.Sync() Fatal.")
		}
	}()

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
