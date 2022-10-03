package board

import (
	"strconv"
	"strings"
)

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
	row, err := strconv.Atoi(move[0])

	if err != nil {
		panic(1)
	}

	column, err := strconv.Atoi(move[1])

	if err != nil {
		panic(1)
	}

	b.board[row][column] = player
}

func NewBoard() Board {
	b := Board{}
	b.Create()

	return b
}
