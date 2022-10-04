package main

import (
	"fmt"
	"os"
	"os/exec"

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

	clear()

	for i := 0; i < maxPlayers; i++ {
		player := players.NewPlayer(os.Stdout, os.Stdin, p.List())

		p.Add(&player)
	}

	for i := 0; i < maxMoves; i++ {
		clear()

		playerTurn := p.Get(0)

		if i%2 != 0 {
			playerTurn = p.Get(1)
		}

		ui.DrawBoard(b)

		fmt.Printf("\n%s move (left|right): ", playerTurn.Name)
		fmt.Scan(&playerTurn.LastMove)

		b.Move(playerTurn.LastMove, playerTurn)
	}
}

func clear() {
	cmd := exec.Command("clear")

	cmd.Stdout = os.Stdout
	cmd.Run()
}
