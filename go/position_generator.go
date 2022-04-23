package giocomani

// PositionGenerator is an iterator that generates
// all of the possible valid game states one at a time.
type PositionGenerator struct {
	hands [4]int
	turn  bool
}

func NewPositionGenerator() PositionGenerator {
	return PositionGenerator{
		hands: [...]int{0, 1, 0, 0},
		turn:  false,
	}
}

func (g *PositionGenerator) Next() (Position, bool) {
	g.hands[3]++

	for i := 3; i >= 1; i-- {
		if g.hands[i] == 5 {
			g.hands[i] = 0
			g.hands[i-1]++
		}
	}

	if g.hands[0] == 5 {
		if g.turn {
			return Position{}, false
		} else {
			*g = NewPositionGenerator()
			g.turn = true
			return g.Next()
		}
	}

	pos := Position{}
	if !g.turn {
		pos.Turn = 0
	} else {
		pos.Turn = 1
	}

	pos.Players[0][0] = g.hands[0]
	pos.Players[0][1] = g.hands[1]
	pos.Players[1][0] = g.hands[2]
	pos.Players[1][1] = g.hands[3]

	return pos, true
}

func (g PositionGenerator) IsLast() bool {
	return g.hands == [...]int{4, 4, 4, 4} && g.turn
}
