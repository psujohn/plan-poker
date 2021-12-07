package game

import (
	"database/sql"
	_ "encoding/json"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setup(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	checkErr(t, err)

	migration, err := os.ReadFile("../db/create_games.sql")
	checkErr(t, err)

	stmt, err := db.Prepare(string(migration))
	checkErr(t, err)

	_, err = stmt.Exec()
	checkErr(t, err)

	return db
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestSave(t *testing.T) {
	db := setup(t)
	defer db.Close()

	tmpGame := NewGame("test")
	id, err := tmpGame.Save(db)
	checkErr(t, err)

	row := db.QueryRow("SELECT * FROM games WHERE id = ? LIMIT 1", id)

	var game Game
	row.Scan(&game.ID, &game.Name)
	if game.Name != "test" {
		t.Errorf("Failed to store game with correct data")
	}
}

func TestFind(t *testing.T) {
	db := setup(t)
	defer db.Close()

	testGame := NewGame("test")
	id, err := testGame.Save(db)
	checkErr(t, err)

	game, err := Find(db, id)
	checkErr(t, err)

	if game.Name != "test" && game.ID != id {
		t.Errorf("Failed to find game")
	}
}
