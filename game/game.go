package game

import (
	"database/sql"

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
