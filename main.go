package main

import (
	"fmt"
	"net/http"
	"plan-poker/game"
)

func main() {
	games := game.NewGames()
	games.AddGame("whisqy")
	games.AddGame("SD")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page")
	})

	mux.HandleFunc("/games", games.Index)
	mux.HandleFunc("/games/", games.Show)
	http.ListenAndServe(":4000", mux)

}
