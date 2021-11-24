package game

import (
	"testing"
)

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
