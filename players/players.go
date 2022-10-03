package players

import (
	"fmt"
	"io"
)

const introduceYourName = "Introduce your name (): "

type Player struct {
	Name   string
	Stamp  string
	Number int
}

type Players struct {
	players []*Player
}

func (p *Players) Add(player *Player) {
	switch len(p.players) {
	case 0:
		player.Stamp = "X"
		player.Number = 1

		p.players = append(p.players, player)
	case 1:
		player.Stamp = "O"
		player.Number = 2

		p.players = append(p.players, player)
	}
}

func (p *Players) Get(number int) *Player {
	return p.List()[number]
}

func (p *Players) List() []*Player {
	return p.players
}

// NewPlayers
func NewPlayer(w io.Writer, r io.Reader) Player {
	player := Player{}

	fmt.Fprint(w, introduceYourName)
	fmt.Fscan(r, &player.Name)

	return player
}
