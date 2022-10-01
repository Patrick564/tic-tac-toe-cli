package main

import (
	"fmt"

	players "github.com/Patrick564/tic-tac-toe-cli/players"
	ui "github.com/Patrick564/tic-tac-toe-cli/ui"
)

func main() {
	players, _ := players.StartNewGame()

	fmt.Println(players)

	ui.DrawBoard()
}
