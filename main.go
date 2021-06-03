package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Welcome to use the simple golang CMS system.")
	http.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})
	fmt.Println("starting server...")
	http.ListenAndServe("127.0.0.1:9999", nil)
}
