package main

import (
	"flag"
	"fmt"
)

// ANSI color code constants for text styling.
const (
	colourBlack  Colour = "\u001b[30m"
	colourRed           = "\u001b[31m"
	colourGreen         = "\u001b[32m"
	colourOrange        = " \033[38;5;208m"
	colourBlue          = "\u001b[34m"
	colourReset         = "\u001b[0m"
)

// Colour type represents ANSI color codes for text styling.
type Colour string

// colorize applies the given ANSI color to the message and prints it to stdout.
func colorize(colour Colour, message string) {
	fmt.Println(colour, message, colourReset)
}

func main() {

	// Define command-line flags
	sky := flag.Bool("sky", false, "display the message")
	blood := flag.Bool("blood", false, "display the message")
	grass := flag.Bool("grass", false, "display the message")
	carrot := flag.Bool("carrot", false, "display the message")
	crow := flag.Bool("crow", false, "display the message")
	flag.Parse()

	// Handle command-line flags
	switch {
	case *sky:
		colorize(colourBlue, "sky is blue")
	case *blood:
		colorize(colourRed, "blood is red")
	case *grass:
		colorize(colourGreen, "grass is green")
	case *carrot:
		colorize(colourOrange, "carrot is orange")
	case *crow:
		colorize(colourBlack, "crow is black")
	default:
		fmt.Println("invalid command!")
	}
}
