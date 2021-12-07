package game

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Game struct {
	ID      int64
	Name    string
	Stories []*Game
}

func NewGame(n string) Game {
	return Game{
		Name: n,
	}
}

func (g *Game) Json() ([]byte, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (g *Game) Save(db *sql.DB) (int64, error) {
	insert := "INSERT INTO games(name) VALUES (?)"
	stmt, err := db.Prepare(insert)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(g.Name)
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	g.ID = id
	return id, nil
}

func Find(db *sql.DB, id int64) (*Game, error) {
	row := db.QueryRow("SELECT * FROM games WHERE id = ? LIMIT 1", id)

	var game Game
	err := row.Scan(&game.ID, &game.Name)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func All(db *sql.DB) ([]*Game, error) {
	rows, err := db.Query("SELECT * FROM games")
	if err != nil {
		log.Printf("Failed to retrieve rows:\n %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var games []*Game
	for rows.Next() {
		var game Game
		rows.Scan(&game.ID, &game.Name)
		games = append(games, &game)
	}

	return games, nil
}
