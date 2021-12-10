package game

import (
	"database/sql"
	_ "encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func prepareFromFile(db *sql.DB, path string) (*sql.Stmt, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(string(data))
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func executeFiles(db *sql.DB, paths ...string) error {
	for _, path := range paths {
		stmt, err := prepareFromFile(db, path)
		if err != nil {
			return err
		}

		_, err = stmt.Exec()
		stmt.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func dbSetup(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	checkErr(t, err)

	err = executeFiles(db, "../db/create_games.sql", "../db/seed_games_sd.sql", "../db/seed_games_whisqy.sql")
	checkErr(t, err)

	return db
}

func TestAll(t *testing.T) {
	db := dbSetup(t)
	defer db.Close()

	games, err := All(db)
	if err != nil {
		t.Fatal(err)
	}

	if count := len(games); count < 2 {
		t.Errorf("Unexpected games count: expected %d got %v", 2, count)
	}

	names := make(map[string]int)
	for _, game := range games {
		names[game.Name] = 1
	}

	if _, found := names["sd"]; !found {
		t.Errorf("Game not found: expected 'sd'")
	}

	if _, found := names["whisqy"]; !found {
		t.Errorf("Game not found: expected 'whisqy'")
	}
}

func TestFindGame(t *testing.T) {
	games := NewGames(nil)
	games.AddGame("test")

	game, err := games.findGame(1)
	if err != nil {
		t.Errorf("Failed to retrieve game: error returned")
	}

	if game.Name != "test" {
		t.Errorf("Failed to retrieve correct game")
	}
}

func TestAddGame(t *testing.T) {
	games := NewGames(nil)
	games.AddGame("test")

	if len(games.games) != 1 {
		t.Errorf("Failed to add game")
	}

	if game := games.games[0]; game.Name != "test" {
		t.Errorf("Failed to AddGame with correct data")
	}
}

func TestShow(t *testing.T) {
	games := NewGames(nil)
	games.AddGame("test")

	// TODO: This is actually borked. The ID is persistent across tests., track ID for game and set here
	req := httptest.NewRequest("GET", "/games/1", nil)
	w := httptest.NewRecorder()
	games.Show(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v expected %v", status, http.StatusOK)
	}
}
