package game

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestFindGame(t *testing.T) {
	games := NewGames()
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
	games := NewGames()
	id := games.AddGame("test")

	if len(games.games) != 1 {
		t.Errorf("Failed to add game")
	}

	game := games.games[0]
	if game.Name != "test" {
		t.Errorf("Failed to AddGame with correct data")
	}
	if game.ID != id {
		t.Errorf("Failed to set ID for new Game")
	}
}

func TestShow(t *testing.T) {
	games := NewGames()
	path := fmt.Sprintf("/games/%d", games.AddGame("test"))

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/games/{id}", games.Show)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v expected %v", status, http.StatusOK)
	}

	var game Game
	if json.Unmarshal(rr.Body.Bytes(), &game); game.Name != "test" {
		t.Errorf("Did not retrieve correct record")
	}
}
