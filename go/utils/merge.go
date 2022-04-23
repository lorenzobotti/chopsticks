package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func MergeLines(dist int, elems ...string) string {
	split := [][]string{}
	for _, elem := range elems {
		lines := strings.Split(elem, "\n")

		split = append(split, lines)
	}

	distInt := strconv.Itoa(dist)
	format := "%-" + distInt + "s"

	out := strings.Builder{}
	lineNum := 0

	for {
		lines := []string{}
		foundALine := false
		for _, elem := range split {
			if len(elem) <= lineNum {
				lines = append(lines, "")
			} else {
				foundALine = true
				lines = append(lines, elem[lineNum])
			}
		}

		if !foundALine {
			break
		}

		for _, line := range lines {
			fmt.Fprintf(&out, format, line)
		}
		fmt.Fprint(&out, "\n")

		lineNum++
	}

	return out.String()
}
