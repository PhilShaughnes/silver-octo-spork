package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PhilShaughnes/siler-octo-spork/internal/web"
)

const (
	port = 5432
)

type Router struct {
	mux *http.ServeMux
}

func (r *Router) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.mux.HandleFunc("GET "+path, handler)
}

func main() {
	fmt.Println(hi("World"))
	r1 := web.NewRouter()
	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", r1))
	v1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("this is octo spork"))
	})

	middleware := web.Use(web.Logging)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: middleware(v1),
	}
	fmt.Println("hello world")

	mux := http.NewServeMux()
	r := *&Router{
		mux: mux,
	}
	// mux.HandleFunc("GET /try/{cmd}", func(w http.ResponseWriter, r *http.Request) {
	r.Get("/try/{cmd}", func(w http.ResponseWriter, r *http.Request) {
		cmd := r.PathValue("cmd")
		if cmd == "panic" {
			panic("beep boop bop")
		}
		fmt.Fprintf(w, "got cmd: %v", cmd)
	})

	v1.Handle("/wrap/", http.StripPrefix("/wrap", r.mux))

	fmt.Printf("Listening on: %v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("uh oh!")
	}

	fmt.Println(bye("for now!"))
}
