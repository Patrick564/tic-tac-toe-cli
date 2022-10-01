package players_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	players "github.com/Patrick564/tic-tac-toe-cli/players"
)

func TestCreatePlayer(t *testing.T) {
	buffer := bytes.Buffer{}
	p := players.Players{}

	firstPlayer := strings.NewReader("Yuuta")
	secondPlayer := strings.NewReader("Rika")

	p.CreatePlayer(&buffer, firstPlayer, players.X)
	p.CreatePlayer(&buffer, secondPlayer, players.O)

	got := p.ListAll()
	want := []players.Player{
		{Name: "Yuuta", Stamp: "X"},
		{Name: "Rika", Stamp: "O"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestPlayerMove(t *testing.T) {}
