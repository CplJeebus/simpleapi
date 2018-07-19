package main

import (
	"fmt"
	"net/http"
	"strings"
)

type sayHello string
type different string

func (g sayHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = string(g)
	w.Write([]byte(message))
}

func (g different) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = string(g)
	w.Write([]byte(message))
}

func main() {
	h := http.NewServeMux()
	fmt.Printf("Started")
	h.Handle("/", sayHello("I'm an API"))
	h.Handle("/api", different("different"))

	if err := http.ListenAndServe(":8090", h); err != nil {
		panic(err)

	}
}
