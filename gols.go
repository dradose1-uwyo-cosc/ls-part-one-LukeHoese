package main

import (
	"os"

	"ls-part-one/functions"
)

func main() {
	args := os.Args[1:]

	// only use color when stdout is a terminal
	useColor := functions.IsTerminal(os.Stdout)

	hadErr := functions.SimpleLS(os.Stdout, args, useColor)
	if hadErr {
		os.Exit(2)
	}
}
