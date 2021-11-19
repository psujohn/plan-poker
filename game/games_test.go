package game

import (
	_ "encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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
	games.AddGame("test")

	// TODO: This is actually borked. The ID is persistent across tests., track ID for game and set here
	req := httptest.NewRequest("GET", "/games/1", nil)
	w := httptest.NewRecorder()
	games.Show(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v expected %v", status, http.StatusOK)
	}
}
