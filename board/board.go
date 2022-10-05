package board

import (
	"errors"
	"strconv"
	"strings"

	players "github.com/Patrick564/tic-tac-toe-cli/players"
)

type Board struct {
	board [][3]string
}

// Create a new 3x3 board with space values " ".
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
func (b *Board) Move(playerMove string, p *players.Player) error {
	move := strings.Split(playerMove, "")

	row, err := strconv.Atoi(move[0])
	if err != nil {
		return err
	}

	column, err := strconv.Atoi(move[1])
	if err != nil {
		return err
	}

	if row > 2 || column > 2 {
		return errors.New("max number is 2")
	}

	if b.board[row][column] != " " {
		return errors.New("box already use")
	}

	b.board[row][column] = p.Stamp

	return nil
}

func (b *Board) Win() bool {
	board := b.Show()

	for _, row := range board {
		if row[0] == row[1] && row[0] == row[2] && row[0] != " " {
			return true
		}
	}

	for idx, column := range board[0] {
		if column == board[1][idx] && column == board[2][idx] && column != " " {
			return true
		}
	}

	if board[0][0] == board[1][1] && board[0][0] == board[2][2] && board[0][0] != " " {
		return true
	}

	if board[0][2] == board[1][1] && board[0][2] == board[2][0] && board[0][2] != " " {
		return true
	}

	return false
}

// Create a new Board instance
func NewBoard() Board {
	b := Board{}
	b.Create()

	return b
}
