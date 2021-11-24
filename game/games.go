package game

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
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

func (g *Games) findGame(id int) (*Game, error) {
	for _, gm := range g.games {
		if gm.ID == id {
			return &gm, nil
		}
	}
	return nil, errors.New("Couldn't find game")
}

func (g *Games) Index(w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal(g.games)
	if err != nil {
		fmt.Println("Error marshaling games data")
	}

	fmt.Fprintf(w, string(payload))
}

func (g *Games) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		fmt.Println("Couldn't parse game ID <", vars["id"], "> in show")
	}

	game, err := g.findGame(int(id))
	if err != nil {
		fmt.Println(err.Error())
	}

	payload, err := json.Marshal(game)
	if err != nil {
		fmt.Println("Error marhaling game data")
	}

	fmt.Fprintf(w, string(payload))
}
