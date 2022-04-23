package giocomani

import "fmt"

type Player [2]int

func (p Player) String() string {
	return fmt.Sprintf("%d %d", p[0], p[1])
}

// CheckHands makes sure that a player's hands can
// never count past 5
func (p *Player) CheckHands() {
	for i, hand := range p {
		if hand >= 5 {
			p[i] = hand % 5
		}
	}
}

// Lost tells you whether a player has
// lost the game (both fists are closed)
func (p Player) Lost() bool {
	return p[0] == 0 && p[1] == 0
}

// CanSplit tells you whether a player can
// "split" their open hand to revive
// their dead one
func (p Player) CanSplit() bool {
	return (p[0] == 0 && p[1]%2 == 0) || (p[1] == 0 && p[0]%2 == 0) && !p.Lost()
	// return false
}

// splits a hand into two if one of the
// player's hands is empty and the other
// one is an even number
func (p *Player) Split() error {
	if !p.CanSplit() {
		return ErrCantSplit{}
	}

	half := p[1] / 2
	if p[1] == 0 {
		half = p[0] / 2
	}

	p[0] = half
	p[1] = half

	return nil
}
