package web

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /hello/{name}", hello)
	router.HandleFunc("GET /snippet/{id}", show)
	router.HandleFunc("GET /snippet", show)
	router.HandleFunc("GET /snippet/create", create)
	return router
}

func hello(w http.ResponseWriter, r *http.Request) {
	v := r.PathValue("name")
	w.Write([]byte("hello " + v))
}

func show(w http.ResponseWriter, r *http.Request) {
	v := r.PathValue("id")
	w.Write([]byte("snippet: " + v))
}

func create(w http.ResponseWriter, _ *http.Request) {
	// v := r.PathValue("name")
	w.Write([]byte("creating snippet"))
}
