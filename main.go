package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = 5432
)

func main() {
	fmt.Println(hi("World"))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /try/{cmd}", func(w http.ResponseWriter, r *http.Request) {
		cmd := r.PathValue("cmd")
		if cmd == "panic" {
			panic("beep boop bop")
		}
		fmt.Fprintf(w, "got cmd: %v", cmd)
	})

	fmt.Printf("Listening on: %v", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux); err != nil {
		log.Fatalf("uh oh!")
	}

	fmt.Println(bye("for now!"))
}
