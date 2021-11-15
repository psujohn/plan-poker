package game

import (
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
