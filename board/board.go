package board

import "strings"

type Board struct {
	board [][3]int
}

// var columns = map[string]int{
// 	"A": 0,
// 	"B": 1,
// 	"C": 2,
// }

// [
// [ 0, 0, 0 ]
// [ 0, 0, 0 ]
// [ 0, 0, 0 ]
// ]
// 0 -> unset
// 1 -> player 1
// 2 -> player 2
// var boardPositions = make([][]int, 3)

func (b *Board) CreateBoard() {
	b.board = make([][3]int, 3)
}

func (b *Board) GetBoard() [][3]int {
	return b.board
}

func PlayerMove(coordinates string) {
	a := strings.Split(coordinates, "")
}

// func InitBoard(coordinateA string, coordinateB int) {}

func ClearBoard() {}
