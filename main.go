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

// FlagInfo struct holds the information about each flag.
type FlagInfo struct {
	flag    *bool
	message string
	colour  Colour
}

// Implement flags as map and include the details
var flags = map[string]FlagInfo{
	"sky":    {nil, "sky is blue", colourBlue},
	"blood":  {nil, "blood is red", colourRed},
	"grass":  {nil, "grass is green", colourGreen},
	"carrot": {nil, "carrot is orange", colourOrange},
	"crow":   {nil, "crow is black", colourBlack},
}

// Set up flags based on the flags map
func initFlags() {
	for key := range flags {
		flag := flag.Bool(key, false, "display the message")
		flags[key] = FlagInfo{flag, flags[key].message, flags[key].colour}
	}
}

// colorize applies the given ANSI color to the message and prints it to stdout.
func colorize(colour Colour, message string) {
	fmt.Println(colour, message, colourReset)
}


// displayMessages checks the flags and prints the message.
func displayMessages() bool {
	displayed := false
	for _, v := range flags {
		if *v.flag {
			colorize(v.colour, v.message)
			displayed = true
		}
	}
	return displayed
}

// generateOptionsList generates a comma-separated list of available options
func generateOptionsList() string {
	options := ""
	for key := range flags {
		if options != "" {
			options += ", "
		}
		options += key
	}
	return options
}


func main() {
	// Print possible options
	fmt.Printf("Possible options for you to flag: [%s]\n", generateOptionsList())

	// Initialize and parse flags
	initFlags()
	flag.Parse()

	// Display messages based on the flags set
	if !displayMessages() {
		fmt.Println("invalid command! Use -h for help.")
	}
}