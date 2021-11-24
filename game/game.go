package game

type Game struct {
	ID      int64
	Name    string
	Stories []*Game
}

func NewGame(id int64, n string) Game {
	return Game{
		ID:   id,
		Name: n,
	}
}
