package game

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func findGame(db *sql.DB, id int64) (*Game, error) {
	row := db.QueryRow("SELECT * FROM games WHERE id = ? LIMIT 1", id)

	var game Game
	err := row.Scan(&game.ID, &game.Name)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func allGames(db *sql.DB) ([]Game, error) {
	rows, err := db.Query("SELECT * FROM games")
	if err != nil {
		log.Printf("Failed to retrieve rows:\n %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var games []Game
	for rows.Next() {
		var game Game
		rows.Scan(&game.ID, &game.Name)
		games = append(games, game)
	}

	return games, nil
}

func IndexHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	f := func(w http.ResponseWriter, r *http.Request) {
		games, err := allGames(db)
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

func ShowHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 32)
		if err != nil {
			log.Printf("Couldn't parse game ID <%s> in show\n", vars["id"])
		}

		game, err := findGame(db, id)
		if err != nil {
			log.Println(err.Error())
		}

		payload, err := json.Marshal(game)
		if err != nil {
			log.Println("Marshal error\n", err.Error())
		}

		fmt.Fprintf(w, string(payload))
	}

	return f
}
