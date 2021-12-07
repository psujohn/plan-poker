package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"plan-poker/game"
	"strconv"

	"github.com/gorilla/mux"
)

func (s server) routes() {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page")
	})

	s.mux.HandleFunc("/games", indexHandler(s.db, game.All))
	s.mux.HandleFunc("/games/{id}", showHandler(s.db, game.Find))
}

func indexHandler(db *sql.DB, all func(*sql.DB) ([]*game.Game, error)) func(http.ResponseWriter, *http.Request) {
	f := func(w http.ResponseWriter, r *http.Request) {
		games, err := all(db)
		if err != nil {
			log.Println(err)
		}

		payload, err := json.Marshal(games)
		if err != nil {
			log.Printf("Error mashaling games data:\n %v\n", err)
		}

		fmt.Fprintf(w, string(payload))
	}

	return f
}

func showHandler(db *sql.DB, find func(*sql.DB, int64) (*game.Game, error)) func(http.ResponseWriter, *http.Request) {
	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 32)
		if err != nil {
			log.Printf("Couldn't parse game ID <%s> in show\n", vars["id"])
		}

		record, err := find(db, id)
		if err != nil {
			log.Println(err.Error())
		}

		payload, err := record.Json()
		if err != nil {
			log.Println("Marshal error\n", err.Error())
		}

		fmt.Fprintf(w, string(payload))
	}

	return f
}
