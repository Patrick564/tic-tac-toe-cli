package main

import (
	"fmt"
	"os"

	board "github.com/Patrick564/tic-tac-toe-cli/board"
	players "github.com/Patrick564/tic-tac-toe-cli/players"
	ui "github.com/Patrick564/tic-tac-toe-cli/ui"
)

const (
	maxMoves   = 9
	maxPlayers = 2
)

func main() {
	p := players.Players{}
	b := board.NewBoard()

	for i := 0; i < maxPlayers; i++ {
		player := players.NewPlayer(os.Stdout, os.Stdin)

		p.Add(&player)
	}

	for i := 0; i < maxMoves; i++ {
		var move string

		fmt.Print("Coordinates (): ")
		fmt.Scan(&move)

		if i == 0 || i%2 == 0 {
			b.Move(move, p.Get(0))
		} else {
			b.Move(move, p.Get(1))
		}

		ui.DrawBoard(b)
	}
}
