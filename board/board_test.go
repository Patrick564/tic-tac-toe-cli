package board_test

import (
	"reflect"
	"testing"

	board "github.com/Patrick564/tic-tac-toe-cli/board"
)

func TestNewBoard(t *testing.T) {
	b := board.NewBoard()

	got := b.Show()
	want := make([][3]string, 3)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %q and got %q", b, want)
	}
}
