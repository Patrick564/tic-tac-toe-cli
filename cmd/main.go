package main

import (
	"fmt"
	"os"

	board "github.com/Patrick564/tic-tac-toe-cli/board"
	players "github.com/Patrick564/tic-tac-toe-cli/players"
	ui "github.com/Patrick564/tic-tac-toe-cli/ui"
)

const maxMoves = 9

func main() {
	p := players.Players{}
	b := board.NewBoard()

	player1 := players.NewPlayer(os.Stdout, os.Stdin)
	player2 := players.NewPlayer(os.Stdout, os.Stdin)

	p.Add(&player1)
	p.Add(&player2)

	for i := 0; i < maxMoves; i++ {
		var m string
		fmt.Print("Introduce your move: ")
		fmt.Scan(&m)

		b.Move(m, player1.Number)
		ui.DrawBoard(b)
	}
}
