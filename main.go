package main

import (
	"fmt"
	"net/http"

	"github.com/nekochans/kimono-app-api/infrastructure/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, HTTPサーバ")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handler)
	http.ListenAndServe(":8888", httputil.Log(router))
}
