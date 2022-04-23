package giocomani

import (
	"errors"
	"fmt"
	"strings"
)

type Move struct {
	Player int
	Split  bool

	From int
	To   int
}

func (m Move) String() string {
	directions := map[int]string{
		0: "l",
		1: "r",
	}

	if m.Split {
		return "split"
	} else {
		return fmt.Sprint(directions[m.From], "-", directions[m.To])
	}
}

func (p Position) Moves() []Move {
	movingPlayer := p.Players[p.Turn]
	otherPlayer := p.Players[1-p.Turn]

	moves := make([]Move, 0, 5)

	if movingPlayer.CanSplit() {
		moves = append(moves, Move{Player: p.Turn, Split: true})
	}

	for myHand, myValue := range movingPlayer {
		if myValue == 0 {
			continue
		}

		for otherHand, otherValue := range otherPlayer {
			if otherValue == 0 {
				continue
			}

			moves = append(moves, Move{
				Player: p.Turn,
				From:   myHand,
				To:     otherHand,
			})
		}
	}

	return moves
}

func (p *Position) Move(m Move) error {
	if p.Turn != m.Player {
		return ErrWrongTurn(m.Player)
	}

	movingPlayer := &p.Players[p.Turn]
	otherPlayer := &p.Players[1-p.Turn]
	if movingPlayer.Lost() {
		return ErrLost{}
	}

	if m.Split {
		err := movingPlayer.Split()
		if err != nil {
			return err
		}
	} else {
		if movingPlayer[m.From] == 0 {
			return ErrEmptyHand(m.From)
		}

		if otherPlayer[m.To] == 0 {
			return ErrEmptyHand(m.To)
		}

		otherPlayer[m.To] += movingPlayer[m.From]
	}

	otherPlayer.CheckHands()
	p.Turn = 1 - p.Turn
	return nil
}

func ParseMove(s string) (Move, error) {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if strings.HasPrefix("split", s) {
		return Move{
			Split: true,
		}, nil
	}

	fromRaw, toRaw, found := strings.Cut(s, "-")
	if !found {
		fromRaw, toRaw, found = strings.Cut(s, "/")
		if !found {
			return Move{}, errors.New("can't parse move: split")
		}
	}

	from := 0
	if fromRaw == "0" || strings.HasPrefix("left", fromRaw) {
		from = 0
	} else if fromRaw == "1" || strings.HasPrefix("right", fromRaw) {
		from = 1
	} else {
		return Move{}, errors.New("can't parse move: from")
	}

	to := 0
	if toRaw == "0" || strings.HasPrefix("left", toRaw) {
		to = 0
	} else if toRaw == "1" || strings.HasPrefix("right", toRaw) {
		to = 1
	} else {
		return Move{}, errors.New("can't parse move: to")
	}

	return Move{
		From: from,
		To:   to,
	}, nil
}
