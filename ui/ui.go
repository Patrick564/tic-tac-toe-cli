package ui

import (
	"fmt"

	board "github.com/Patrick564/tic-tac-toe-cli/board"
	players "github.com/Patrick564/tic-tac-toe-cli/players"
)

func DrawBoard(board board.Board) {
	for idx, row := range board.Show() {
		if idx == 0 {
			fmt.Print("\n")
			fmt.Println(`   0   1   2`)
			fmt.Println(` +-----------`)
		}

		fmt.Printf("%d| %s | %s | %s \n", idx, row[0], row[1], row[2])

		if idx < 2 {
			fmt.Println(` |---|---|---`)
		}
	}
}

func DrawWinner(p players.Player) {
	fmt.Printf("\nPlayer %s is the winner!", p.Name)
}
