package main

import "github.com/nekochans/kimono-app-api/infrastructure"

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello, HTTPサーバ")
//}

func main() {
	infrastructure.StartHTTPServer()
	//router := http.NewServeMux()
	//router.HandleFunc("/", handler)
	//http.ListenAndServe(":8888", httputil.Log(router))
}
