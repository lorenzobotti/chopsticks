package utils

import (
	"fmt"
	"os"
)

func Handle(err any) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
