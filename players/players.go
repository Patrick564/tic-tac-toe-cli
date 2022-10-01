package players

import (
	"fmt"
	"io"
)

const (
	introduceYourName = "Introduce your nickname: "
	selectYourStamp   = "Select your stamp ('X' or 'O'): "
)

const (
	X = "X"
	O = "O"
)

type Player struct {
	Name  string
	Stamp string
}

type Players struct {
	players []Player
}

type Board struct{}

func (p *Players) CreatePlayer(w io.Writer, r io.Reader, stamp string) {
	player := Player{Stamp: stamp}

	fmt.Fprint(w, introduceYourName)
	fmt.Fscan(r, &player.Name)

	p.players = append(p.players, player)
}

func (p *Players) ListAll() []Player {
	return p.players
}

func PlayerMove(coordinateA string, coordinateB int) {}
