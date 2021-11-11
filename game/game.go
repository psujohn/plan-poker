package game

var lastId = 0

type Game struct {
	ID   int
	Name string
}

func NewGame(n string) *Game {
	lastId++

	return &Game{
		ID:   lastId,
		Name: n,
	}
}
