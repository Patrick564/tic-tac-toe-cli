package board

import (
	"strconv"
	"strings"

	"github.com/Patrick564/tic-tac-toe-cli/players"
)

type Board struct {
	board [][3]string
}

// Create a new board with space values (" ")
func (b *Board) Create() {
	b.board = make([][3]string, 3)

	for idx, row := range b.board {
		for box := range row {
			b.board[idx][box] = " "
		}
	}
}

func (b *Board) Reset() {
	b.board = make([][3]string, 0)
}

func (b *Board) Show() [][3]string {
	return b.board
}

// Set a player move using coordinates (y, x) and player instance
func (b *Board) Move(coordinates string, player *players.Player) {
	move := strings.Split(coordinates, "")
	row, err := strconv.Atoi(move[0])

	if err != nil {
		panic(1)
	}

	column, err := strconv.Atoi(move[1])

	if err != nil {
		panic(1)
	}

	b.board[row][column] = player.Stamp
}

// Create a new Board instance
func NewBoard() Board {
	b := Board{}
	b.Create()

	return b
}
