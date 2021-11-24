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

type Games struct {
	games []Game
	seq   int
}

func NewGames() Games {
	return Games{seq: 0}
}

func AddGame(db *sql.DB, name string) (int64, error) {
	insert := "INSERT INTO games(name) VALUES (?)"
	stmt, err := db.Prepare(insert)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name)
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

func findGame(db *sql.DB, id int64) (*Game, error) {
	row := db.QueryRow("SELECT * FROM games WHERE id = ? LIMIT 1", id)

	var game Game
	err := row.Scan(&game.ID, &game.Name)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func (g *Games) Index(w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal(g.games)
	if err != nil {
		fmt.Println("Error marshaling games data")
	}

	fmt.Fprintf(w, string(payload))
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
