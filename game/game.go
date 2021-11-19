package game

type Game struct {
	ID      int
	Name    string
	Stories []*Game
}

func NewGame(id int, n string) Game {
	return Game{
		ID:   id,
		Name: n,
	}
}
