package board_test

import (
	"reflect"
	"testing"

	board "github.com/Patrick564/tic-tac-toe-cli/board"
)

func TestCreateBoard(t *testing.T) {
	b := board.Board{}
	b.CreateBoard()

	got := b.GetBoard()
	want := make([][3]int, 3)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %q and got %q", b, want)
	}
}
