package ui

import (
	"fmt"

	board "github.com/Patrick564/tic-tac-toe-cli/board"
)

func DrawBoard(board board.Board) {
	for idx, row := range board.Show() {
		if idx == 0 {
			fmt.Println(`   0   1   2`)
			fmt.Println(` +-----------`)
		}

		fmt.Printf("%d| %s | %s | %s \n", idx, row[0], row[1], row[2])

		if idx < 2 {
			fmt.Println(` |---|---|---`)
		}
	}
}
