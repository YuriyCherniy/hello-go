package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, `<!DOCTYPE html><html><head><title>Hello</title></head><body>hello world</body></html>`)
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server at :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
