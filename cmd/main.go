package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

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

		err := b.Move(playerTurn.LastMove, playerTurn)

		if err != nil {
			fmt.Print("Only numbers from 0 to 2 permitted.")
			time.Sleep(time.Second)
			i--

			continue
		}
	}
}

func clear() {
	cmd := exec.Command("clear")

	cmd.Stdout = os.Stdout
	cmd.Run()
}
