package main

import (
	"fmt"
	"net/http"
	"plan-poker/game"

	"github.com/gorilla/mux"
)

func main() {
	games := game.NewGames()
	games.AddGame("whisqy")
	games.AddGame("SD")

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page")
	})

	r.HandleFunc("/games", games.Index)
	r.HandleFunc("/games/{id}", games.Show)

	http.ListenAndServe(":4000", r)
}
