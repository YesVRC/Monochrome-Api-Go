package main

import (
	"fmt"
	"net/http"
)

func main() {
	s := http.NewServeMux()

	s.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "Hello World")
	})
	err := http.ListenAndServe(":8080", s)
	if err != nil {
		panic(err)
	}
}
