package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"plan-poker/game"
	"strconv"
)

var games []*game.Game

func findGame(id int) *game.Game {
	for _, g := range games {
		if g.ID == id {
			return g
		}
	}
	return nil
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.URL.Path[len("/games/"):], 10, 32)
	g := findGame(int(id))
	s, _ := json.Marshal(g)
	fmt.Fprintf(w, string(s))
}

func main() {
	games = append(games, game.NewGame("whisqy"))
	games = append(games, game.NewGame("SD"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page")
	})
	mux.HandleFunc("/games", func(w http.ResponseWriter, req *http.Request) {
		s, _ := json.Marshal(games)
		fmt.Fprintf(w, string(s))
	})
	mux.HandleFunc("/games/", gameHandler)
	http.ListenAndServe(":4000", mux)

}
