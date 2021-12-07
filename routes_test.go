package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"plan-poker/game"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func find(db *sql.DB, id int64) (*game.Game, error) {
	game := &game.Game{
		ID:   id,
		Name: "test name",
	}

	return game, nil
}

func all(db *sql.DB) ([]*game.Game, error) {
	games := []*game.Game{
		{ID: 13, Name: "test name"},
	}

	return games, nil
}

func TestShowHandler(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/games/13", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/games/{id}", showHandler(db, find))
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code:  got %v expected %v", status, http.StatusOK)
	}

	var g game.Game
	err = json.Unmarshal(rr.Body.Bytes(), &g)
	if err != nil {
		t.Fatal(err)
	}

	if g.ID != 13 {
		t.Errorf("Wrong ID: got %v expected %v", g.ID, 13)
	}
}

func TestIndexHandler(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/games", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler(db, all))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v expected %v", status, http.StatusOK)
	}

	var gs []game.Game
	err = json.Unmarshal(rr.Body.Bytes(), &gs)
	if err != nil {
		t.Fatal(err)
	}

	if gs[0].Name != "test name" {
		t.Errorf("Wrong record(s) retrieved: got %v expected %v", gs[0].Name, "test name")
	}
}
