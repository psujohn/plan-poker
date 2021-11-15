package game

import (
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
)

type Games struct {
  games []Game
  seq int
}

func NewGames() Games {
  return Games { seq: 0 }
}

func (g *Games) AddGame(name string) {
  g.games = append(g.games, NewGame(name))
}

func (g Games) findGame(id int) *Game{
  for _, gm := range g.games {
    if gm.ID == id {
      return &gm
    }
  }
  return nil
}

func (g Games) Index(w http.ResponseWriter, r *http.Request) {
  payload, err := json.Marshal(g.games)
  if err != nil {
    fmt.Println("Error marshaling games data")
  }

  fmt.Fprintf(w, string(payload))
}

func (g Games) Show(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.ParseInt(r.URL.Path[len("/games/"):], 10, 32)
  if err != nil {
    fmt.Println("Couldn't parse game ID in show")
  }

  game := g.findGame(int(id))
  payload, err := json.Marshal(game)
  if err != nil {
    fmt.Println("Error marhaling game data")
  }

  fmt.Fprintf(w, string(payload))
}
