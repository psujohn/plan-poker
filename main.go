package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"plan-poker/game"
)

func main() {
	var games []*game.Game

	games = append(games, game.NewGame("whisqy"))
	games = append(games, game.NewGame("SD"))

	s, _ := json.Marshal(games)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page")
	})
	mux.HandleFunc("/games", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, string(s))
	})
	http.ListenAndServe(":4000", mux)

}
