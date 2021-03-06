package main

import (
	"fmt"
	"net/http"
)

func (s server) routes() {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page")
	})

	s.mux.HandleFunc("/games", s.games.Index)
	s.mux.HandleFunc("/games/{id}", s.games.Show)
}
