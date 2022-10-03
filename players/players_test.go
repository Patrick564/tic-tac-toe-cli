package players_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	players "github.com/Patrick564/tic-tac-toe-cli/players"
)

func TestCreate(t *testing.T) {
	buffer := bytes.Buffer{}
	p := players.Players{}

	name1 := strings.NewReader("Yuuta")
	name2 := strings.NewReader("Rika")

	player1 := players.NewPlayer(&buffer, name1)
	player2 := players.NewPlayer(&buffer, name2)

	p.Add(player1)
	p.Add(player2)

	got := p.List()
	want := []players.Player{
		{Name: "Yuuta", Stamp: "X", Number: 1},
		{Name: "Rika", Stamp: "O", Number: 2},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
