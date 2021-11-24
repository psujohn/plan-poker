package game

import (
	"database/sql"
	_ "encoding/json"
	_ "net/http"
	_ "net/http/httptest"
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

func TestFindGame(t *testing.T) {
}

func TestAddGame(t *testing.T) {
	db := setup(t)
	defer db.Close()

	id, err := AddGame(db, "test")
	checkErr(t, err)

	row, err := db.Query("SELECT * FROM games WHERE id = ? LIMIT 1", id)
	checkErr(t, err)
	defer row.Close()

	var game Game
	for row.Next() {
		row.Scan(&game.ID, &game.Name)
	}
	if game.Name != "test" {
		t.Errorf("Failed to store game with correct data")
	}
}

func TestShow(t *testing.T) {
	/*
	  games := NewGames()
		games.AddGame("test")

		// TODO: This is actually borked. The ID is persistent across tests., track ID for game and set here
		req := httptest.NewRequest("GET", "/games/1", nil)
		w := httptest.NewRecorder()
		games.Show(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("Wrong status code: got %v expected %v", status, http.StatusOK)
		}
	*/
}
