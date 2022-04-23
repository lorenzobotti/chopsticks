package giocomani

type Searcher struct {
	Starting    Position
	seenAlready map[Position]int
}

const (
	ScoreDraw = iota - 1
	ScoreOne
	ScoreTwo
)

func oppositeScore(score int) int {
	switch score {
	case ScoreDraw:
		return ScoreDraw
	case ScoreOne:
		return ScoreTwo
	case ScoreTwo:
		return ScoreOne
	default:
		panic("unimplemented!")
	}
}

func ScoreString(score int) string {
	switch score {
	case ScoreDraw:
		return "draw"
	case ScoreOne:
		return "player one"
	case ScoreTwo:
		return "player two"
	default:
		panic("unimplemented!")
	}
}

func NewSearcher() Searcher {
	s := Searcher{}
	s.seenAlready = map[Position]int{}
	return s
}

func (s *Searcher) BestMove(pos Position, ply int) (Move, int) {
	moves := pos.Moves()
	bestScore := oppositeScore(pos.Turn)
	bestMove := moves[0]

	for _, move := range moves {
		cloned := pos
		cloned.Move(move)

		s.Starting = cloned
		score := s.Search(ply)

		if score == pos.Turn {
			return move, score
		} else if score == ScoreDraw {
			bestScore = ScoreDraw
			bestMove = move
		}
	}

	return bestMove, bestScore
}

func (s *Searcher) Search(ply int) int {
	return s.searchRec(s.Starting, ply)
}

func (s *Searcher) searchRec(pos Position, ply int) int {
	over, winner := pos.GameOver()
	if over {
		return winner
	}

	// if ply == 0 {
	// 	return ScoreDraw
	// }

	maximisingPlayer := pos.Turn
	moves := pos.Moves()

	bestScore := oppositeScore(maximisingPlayer)
	for _, move := range moves {
		cloned := pos
		cloned.Move(move)

		score, seen := s.seenAlready[cloned]
		if !seen {
			s.seenAlready[cloned] = ScoreDraw
			score = s.searchRec(cloned, ply-1)

			s.seenAlready[cloned] = score
		}

		// fmt.Println(ply, "ply")
		// fmt.Println(cloned)

		if score == maximisingPlayer {
			return score
		}

		if score == ScoreDraw {
			bestScore = ScoreDraw
		}
	}

	return bestScore
}

// func (s Searcher) String() string {
// 	out := strings.Builder{}
// 	posStr := []string{}
// 	for _, pos := range s.pos {
// 		posStr = append(posStr, pos.String())
// 	}

// 	for _, chunk := range utils.Chunk(posStr, 5) {
// 		merged := utils.MergeLines(10, chunk...)
// 		out.WriteString(merged)
// 	}

// 	return out.String()
// }
