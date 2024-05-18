package main

import (
	"flag"
	"fmt"
	"strings"
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

// Implement flags as map and include the details and use FlagInfo as pointer
var flags = map[string]*FlagInfo{
	"sky":    {message: "sky is blue", colour: colourBlue},
	"blood":  {message: "blood is red", colour: colourRed},
	"grass":  {message: "grass is green", colour: colourGreen},
	"carrot": {message: "carrot is orange", colour: colourOrange},
	"crow":   {message: "crow is black", colour: colourBlack},
}

// Set up flags based on the flags map
func initFlags() {
	for key, value := range flags {
		value.flag = flag.Bool(key, false, "display the message")
	}
}

// colorize applies the given ANSI color to the message and prints it to stdout.
func colourize(colour Colour, message string) {
	fmt.Println(colour, message, colourReset)
}

// displayMessages checks the flags and prints the message.
func displayMessages() bool {
	displayed := false
	for _, v := range flags {
		if *v.flag {
			colourize(v.colour, v.message)
			displayed = true
		}
	}
	return displayed
}

// generateOptionsList generates a comma-separated list of available options
func generateOptionsList() string {
	var options []string
	for key := range flags {
		options = append(options, key)
	}
	return strings.Join(options, ", ")
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
