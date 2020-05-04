package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nekochans/kimono-app-api/infrastructure"
)

func handler(w http.ResponseWriter, r *http.Request) {
	logOpts := infrastructure.LoggerOptions{}
	logger, err := infrastructure.NewLoggerFromOptions(logOpts)

	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] Failed to create logger: %s\n", err)
		os.Exit(1)
	}

	defer logger.Sync()
	logger.Info("info test log")

	fmt.Fprint(w, "Hello, HTTPサーバ")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
