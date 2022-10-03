package board

import (
	"strconv"
	"strings"
)

var columns = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
}

type Board struct {
	board [][3]int
}

func (b *Board) Create() {
	b.board = make([][3]int, 3)
}

func (b *Board) Reset() {
	b.board = make([][3]int, 0)
}

func (b *Board) Show() [][3]int {
	return b.board
}

func (b *Board) Move(coordinates string, player int) {
	move := strings.Split(coordinates, "")
	row := columns[move[0]]
	column, _ := strconv.Atoi(move[1])

	b.board[row][column] = player
}

func NewBoard() Board {
	b := Board{}
	b.Create()

	return b
}
