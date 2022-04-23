package giocomani

import (
	"fmt"
	"testing"
)

func TestGenerateMoves(t *testing.T) {
	cases := []struct {
		pos   Position
		moves []Move
	}{
		{
			pos: Position{
				Turn: 0,
				Players: [...]Player{
					{2, 3},
					{1, 4},
				},
			},
			moves: []Move{
				{Player: 0, From: 0, To: 0},
				{Player: 0, From: 0, To: 1},
				{Player: 0, From: 1, To: 0},
				{Player: 0, From: 1, To: 1},
			},
		},
		{
			pos: Position{
				Turn: 0,
				Players: [...]Player{
					{0, 3},
					{1, 4},
				},
			},
			moves: []Move{
				{Player: 0, From: 1, To: 0},
				{Player: 0, From: 1, To: 1},
			},
		},
		{
			pos: Position{
				Turn: 1,
				Players: [...]Player{
					{2, 3},
					{0, 4},
				},
			},
			moves: []Move{
				{Player: 1, Split: true},
				{Player: 1, From: 1, To: 0},
				{Player: 1, From: 1, To: 1},
			},
		},
	}

	for tcI, tc := range cases {
		moves := tc.pos.Moves()
		for i, exp := range tc.moves {
			got := moves[i]

			if exp != got {
				fmt.Println(moves)
				t.Fatalf("case %d, move %d: expected move %v, got move %v", tcI, i, exp, got)
			}
		}
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		pos      Position
		move     Move
		expected Position
		err      error
	}{
		{
			pos: Position{
				Turn: 0,
				Players: [2]Player{
					{1, 3},
					{4, 2},
				},
			},
			move: Move{Player: 0, From: 0, To: 0},
			expected: Position{
				Turn: 1,
				Players: [2]Player{
					{1, 3},
					{0, 2},
				},
			},
			err: nil,
		},
		{
			pos: Position{
				Turn: 0,
				Players: [2]Player{
					{0, 4},
					{4, 2},
				},
			},
			move: Move{Player: 0, Split: true},
			expected: Position{
				Turn: 1,
				Players: [2]Player{
					{2, 2},
					{4, 2},
				},
			},
			err: nil,
		},
		{
			pos: Position{
				Turn: 0,
				Players: [2]Player{
					{0, 3},
					{4, 2},
				},
			},
			move: Move{Player: 0, From: 0, To: 0},
			err:  ErrEmptyHand(0),
		},
	}

	for tcI, tc := range testCases {
		cloned := tc.pos
		err := cloned.Move(tc.move)

		if err != tc.err {
			t.Fatalf("expected error '%v', got '%v", tc.err, err)

		}

		if err != nil {
			continue
		}

		if cloned != tc.expected {
			fmt.Println("case", tcI)
			fmt.Println(cloned)
			t.Fail()
		}
	}
}
