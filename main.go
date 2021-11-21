package main

import (
	"net/http"
	"plan-poker/game"

	"github.com/gorilla/mux"
)

func main() {
	srv, err := newServer()
	if err != nil {
		return
	}
	srv.games.AddGame("whisqy")
	srv.games.AddGame("SD")

	http.ListenAndServe(":4000", srv.mux)
}

type server struct {
	mux   *mux.Router
	games *game.Games
}

func newServer() (*server, error) {
	g := game.NewGames()
	srv := &server{
		mux:   mux.NewRouter(),
		games: &g,
	}
	srv.routes()
	return srv, nil
}
