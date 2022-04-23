package giocomani

import (
	"fmt"
	"strings"
)

type Position struct {
	Turn    int
	Players [2]Player
}

func DefaultPosition() Position {
	return Position{
		Turn: 0,
		Players: [2]Player{
			{1, 1},
			{1, 1},
		},
	}
}

// returns wether the game is over or not and
// who the winner is
func (p Position) GameOver() (bool, int) {
	if p.Players[0].Lost() {
		return true, 1
	} else if p.Players[1].Lost() {
		return true, 0
	} else {
		return false, 0
	}
}

// String formats the position in a human readable way.
// it look like this:
//     > 2 4
//       3 0
func (p Position) String() string {
	build := strings.Builder{}

	over, winner := p.GameOver()

	for player, hand := range p.Players {
		turn := "-"
		if p.Turn == player {
			turn = ">"
		}

		lost := ""
		if over && winner != player {
			lost = ":("
		}

		fmt.Fprintln(&build, turn, hand[0], hand[1], lost)
	}

	return build.String()
}

func (p Position) CompactString() string {
	return fmt.Sprintf(
		"%d -> %d-%d | %d-%d",
		p.Turn,
		p.Players[0][0],
		p.Players[0][1],
		p.Players[1][0],
		p.Players[1][1],
	)
}
