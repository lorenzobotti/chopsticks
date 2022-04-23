package main

import (
	"fmt"
	. "giocomani"
	"giocomani/utils"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if containsArgument("generate") {
		generateMoves(os.Stdout)
	} else {
		repl()
	}
}

func containsArgument(arg string) bool {
	arg = strings.ToLower(arg)
	for _, ar := range os.Args {
		if strings.ToLower(ar) == arg {
			return true
		}
	}

	return false
}

func generateMoves(out io.Writer) {
	gen := NewPositionGenerator()
	searcher := NewSearcher()
	ply := 30

	fmt.Fprintln(out, "[")

	for {
		pos, ok := gen.Next()
		if !ok {
			break
		}

		over, score := pos.GameOver()
		best := Move{}

		if !over {
			best, score = searcher.BestMove(pos, ply)
		}

		fmt.Fprintf(
			out,
			`{"%s": {"move": "%s", "score": "%s", "over": %s, "can_split": %s, "moves": {`,
			pos.CompactString(),
			best.String(),
			ScoreString(score),
			boolString(over),
			boolString(pos.Players[pos.Turn].CanSplit()),
		)

		for _, move := range pos.Moves() {
			cloned := pos
			cloned.Move(move)

			fmt.Fprintf(
				out,
				`"%s": "%s",`,
				move.String(),
				cloned.CompactString(),
			)
		}

		fmt.Fprint(out, "}}}")
	}

	fmt.Fprint(out, "]")
	closer, ok := out.(io.WriteCloser)
	if ok {
		closer.Close()
	}
}

func repl() {
	ply := 30
	pos := DefaultPosition()
	s := NewSearcher()
	aiPlayer := 1

	for {
		over, winner := pos.GameOver()
		if over {
			if winner == aiPlayer {
				fmt.Println("i win, bye bye")
			} else {
				fmt.Println("you win, you defeated me :(")
			}

			break
		}

		fmt.Println(pos)

	moveinput:
		move := Move{}
		fmt.Print("Your move: ")
		moveRaw := ""
		_, err := fmt.Scanln(&moveRaw)
		utils.Handle(err)

		move, err = ParseMove(moveRaw)
		if err != nil {
			fmt.Println("error:", err)
			goto moveinput
		}

		err = pos.Move(move)
		if err != nil {
			fmt.Println("error:", err)
			goto moveinput
		}

		fmt.Println(pos)
		over, winner = pos.GameOver()
		if over {
			if winner == aiPlayer {
				fmt.Println("i win, bye bye")
			} else {
				fmt.Println("you win, you defeated me :(")
			}

			break
		}

		best, score := s.BestMove(pos, ply)
		err = pos.Move(best)
		utils.Handle(err)

		fmt.Println("i play", best)
		fmt.Println("my prediction:", ScoreString(score))
	}
}

func countWins() {
	ply := 30
	search := NewSearcher()
	generator := NewPositionGenerator()
	one, two, draws := 0, 0, 0
	positionsEvaluated := 0
	totalPositions := 6 * 6 * 6 * 6 * 2
	outputFrequency := time.Second

	start := time.Now()
	lastPrinted := time.Now()
	logStatus := func() {
		fmt.Printf(
			"evaluated %10d/%d positions at depth %d. %10d draws, %10d wins for one, %10d wins for two. %f s\n",
			positionsEvaluated,
			totalPositions,
			ply,
			draws, one, two,
			time.Since(start).Seconds(),
		)
	}

	output, err := os.Create("mani_endless.json")
	utils.Handle(err)
	defer output.Close()

	fmt.Fprintln(output, "{")

	for {
		pos, notOver := generator.Next()
		positionsEvaluated++
		if !notOver {
			break
		}

		if time.Since(lastPrinted) > outputFrequency {
			logStatus()
			lastPrinted = time.Now()
		}

		over, _ := pos.GameOver()
		if over {
			continue
		}

		search.Starting = pos
		result := search.Search(ply)
		scoreString := ""

		switch result {
		case ScoreDraw:
			draws++
			scoreString = "draw"
		case ScoreOne:
			one++
			scoreString = "one"
		case ScoreTwo:
			two++
			scoreString = "two"
		}

		fmt.Fprintf(
			output,
			"\t\"%s\": \"%s\",\n",
			pos.CompactString(),
			scoreString,
		)
	}

	fmt.Fprint(output, "}")

	logStatus()
}

func genLoop() {
	gen := PositionGenerator{}

	for {
		numRaw := ""
		_, err := fmt.Scan(&numRaw)
		utils.Handle(err)

		num, err := strconv.Atoi(numRaw)
		utils.Handle(err)

		for i := 0; i < num-1; i++ {
			gen.Next()
		}

		pos, stillGoing := gen.Next()
		if !stillGoing {
			break
		}
		fmt.Println(pos)
	}

	fmt.Println("out of positions")
}

func boolString(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}
