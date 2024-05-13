package main

import (
	"flag"
	"fmt"
)

type Colour string

const (
	colourBlack  Colour = "\u001b[30m"
	colourRed          = "\u001b[31m"
	colourGreen        = "\u001b[32m"
	colourYellow       = "\u001b[33m"
	colourBlue         = "\u001b[34m"
	colourReset        = "\u001b[0m"
)

func colorize(colour Colour, message string) {
	fmt.Println(colour, message, colourReset)
}

func main() {

	sky := flag.Bool("sky", false, "display the message")
	blood := flag.Bool("blood", false, "display the message")
	flag.Parse()

	switch {
	case *sky:
		colorize(colourBlue, "sky is blue")
	case *blood:
		colorize(colourRed, "blood is red")
	default:
		fmt.Println("invalid command!")
	}
}
