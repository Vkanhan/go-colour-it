package main

import (
	"flag"
	"fmt"
)

type Colour string

const (
	colourRed   Colour = "\u001b[31m"
	colourBlue  Colour = "\u001b[34m"
	colourReset Colour = "\u001b[0m"
)

func colorize(colour Colour, message string) {
	fmt.Println(colour, message, colourReset)
}

func main() {

	sky := flag.Bool("sky", false, "display the message")
	blood := flag.Bool("blood", false, "display the message")
	flag.Parse()

	if *sky {
		colorize(colourBlue, "sky is blue")
	}

	if *blood {
		colorize(colourRed, "blood is red")
	}
}
