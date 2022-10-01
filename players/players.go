package players

import (
	"fmt"
)

// [
// [ 0, 0, 0 ]
// [ 0, 0, 0 ]
// [ 0, 0, 0 ]
// ]
// 0 -> unset
// 1 -> player 1
// 2 -> player 2
// var boardPositions = make([][]string, 0)

type Player struct {
	Name  string
	Stamp string
}

type Board struct{}

func StartNewGame() (string, error) {
	n := Player{}

	fmt.Print("Introduce your nickname: ")
	fmt.Scan(&n.Name)

	fmt.Print("Select your stamp ('X' or 'O'): ")
	fmt.Scan(&n.Stamp)

	return fmt.Sprintf("Player 1 (%s): %s", n.Stamp, n.Name), nil
}

func PlayerMove() {}

func ClearBoard() {}
