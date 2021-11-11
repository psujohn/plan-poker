package main

import (
	"fmt"
	"net/http"
	"plan-poker/game"
)

func main() {
	var games []*game.Game

	games = append(games, game.NewGame("whisqy"))
	games = append(games, game.NewGame("SD"))

	/*
		for _, g := range games {
			fmt.Println(*g)
		}
	*/

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page")
	})
	http.ListenAndServe(":4000", mux)

}
