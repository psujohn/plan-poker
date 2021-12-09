package main

import (
	"database/sql"
	"log"
	"net/http"
	"plan-poker/game"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./planning-poker.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	srv, err := newServer(db)
	if err != nil {
		return
	}
	srv.games.AddGame("whisqy")
	srv.games.AddGame("SD")

	http.ListenAndServe(":4000", srv.mux)
}

type server struct {
	db    *sql.DB
	mux   *mux.Router
	games *game.Games
}

func newServer(db *sql.DB) (*server, error) {
	g := game.NewGames()
	srv := &server{
		db:    db,
		mux:   mux.NewRouter(),
		games: &g,
	}
	srv.routes()
	return srv, nil
}
