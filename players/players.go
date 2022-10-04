package players

import (
	"fmt"
	"io"
)

const introduceYourName = "Name (player %d): "

type Player struct {
	Name     string
	Stamp    string
	Number   int
	LastMove string
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

func NewPlayer(w io.Writer, r io.Reader, l []*Player) Player {
	player := Player{}

	fmt.Fprintf(w, introduceYourName, len(l)+1)
	fmt.Fscan(r, &player.Name)

	return player
}
