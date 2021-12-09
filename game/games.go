package game

import (
  "database/sql"
	"encoding/json"
	"errors"
	"fmt"
  "log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
  _ "github.com/mattn/go-sqlite3"
)

type Games struct {
	games []*Game
	seq   int
}

func All(db *sql.DB) ([]*Game, error) {
  rows, err := db.Query("SELECT * FROM games")
  if err != nil {
    log.Printf("Failed to retrieve games:\n %v\n", err)
    return nil, err
  }
  defer rows.Close()

  var games []*Game
  for rows.Next() {
    var game Game
    rows.Scan(&game.ID, &game.Name)
    games = append(games, &game)
  }

  return games,nil
}
func NewGames(db *sql.DB) Games {
  if db == nil {
    return Games{seq: 0}
  }

  games, err := All(db)
  if err != nil {
    return Games{seq: 0}
  }

  return Games{games: games, seq: 0}
}

func (g *Games) AddGame(name string) {
	g.games = append(g.games, NewGame(name))
}

func (g *Games) findGame(id int) (*Game, error) {
	for _, gm := range g.games {
		if gm.ID == id {
			return gm, nil
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
