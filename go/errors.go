package giocomani

import "fmt"

type ErrWrongTurn int
type ErrEmptyHand int
type ErrCantSplit struct{}
type ErrLost struct{}

func (w ErrWrongTurn) Error() string {
	return fmt.Sprint("wrong turn, moving player should be ", 1-w)
}

func (h ErrEmptyHand) Error() string {
	return fmt.Sprint("cannot attack with empty hand ", int(h))
}

func (ErrCantSplit) Error() string {
	return "cannot split without an empty hand or without an even number"
}

func (ErrLost) Error() string {
	return "a player that has lost cannot move"
}
