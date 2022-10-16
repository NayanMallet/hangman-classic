package game

import (
	"fmt"
	g_func "hangman-classic/g-func"
	"hangman-classic/overlay"
	"hangman-classic/presentations"
	"strings"
)

func Progress() {
	// This function will be used to print the progress of the game.

	// Print the rules of the game.
	overlay.Wrapper(overlay.PrintTop, presentations.GamesRules, overlay.PrintBottom)
	presentations.Sleep(1)
	presentations.ClearTerminal()

	overlay.PrintTop()
	presentations.GameTypeChoice()
	presentations.Sleep(1)
	fmt.Printf("Your choice: \n")
	typeChoice := strings.ToUpper(g_func.GetInput())
	switch typeChoice {
	case "LOAD":
		// Load the last save.
		presentations.ClearTerminal()
		overlay.PrintTop()
		presentations.Sleep(1)
		Game("--startWith save.txt")
	case "NEW":
		// start a new game
		presentations.ClearTerminal()
		overlay.PrintTop()
		presentations.GameOptionChoice()
		fmt.Printf("Your choice: \n")
		wordListChoice := strings.ToUpper(g_func.GetInput())

		presentations.ClearTerminal()
		overlay.PrintTop()
		presentations.Sleep(1)
		presentations.GameAsciiOptionChoice()
		fmt.Printf("Your choice: \n")
		asciiChoice := strings.ToUpper(g_func.GetInput())

		presentations.ClearTerminal()
		overlay.PrintTop()
		presentations.Sleep(1)
		if wordListChoice == "N" {
			// start a new game with default word list
			if asciiChoice == "N" {
				// start a new game without ascii art
				Game("words.txt")
			} else {
				// start a new game with custom ascii art
				Game("words.txt --letterFile " + asciiChoice)
			}
		} else {
			// start a new game with custom word list
			if asciiChoice == "N" {
				// start a new game without ascii art
				Game(wordListChoice)
			} else {
				// start a new game with custom ascii art
				Game(wordListChoice + " --letterFile " + asciiChoice)
			}
		}
	}
	overlay.PrintBottom()
}
