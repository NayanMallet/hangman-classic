package presentations

import (
	"fmt"
)

func GamesRules() {
	// print the rules of the hangman game
	rules := []string{
		"Rules are simple, you have to guess the word by suggesting letters.",
		"You can suggest letter and the whole word.",
		"If the letter suggest is correct, the letter will be revealed in the word, else you will lose an attempt.",
		"If your word suggest is correct you win, else you will lose 2 attempts.",
		"If you lose all your attempts, you will lose the game."}
	for i := 0; i < len(rules); i++ {
		fmt.Printf("%s\n", Center(rules[i], "\u0020"))
		Sleep(1)
	}
}

func GameTypeChoice() {
	// print the possibles choices of the hangman game
	fmt.Printf("%s\n", Center("If you want to load last save, type 'LOAD'.", "\u0020"))
	fmt.Printf("%s\n", Center("If you want to start a new game, type 'NEW'.", "\u0020"))
}

func GameOptionChoice() {
	// print the possibles choices of the hangman game
	fmt.Printf("%s\n", Center("For load custom words file enter the file name :\n", "\u0020"))
	fmt.Printf("%s\n", Center("Else tap 'N' :)\n", "\u0020"))
}

func GameAsciiOptionChoice() {
	// print the possibles Ascii Options of the hangman game
	fmt.Printf("%s\n", Center("For load custom ascii art file enter the file name :\n", "\u0020"))
	fmt.Printf("%s\n", Center("Else tap 'N' :)\n", "\u0020"))
}
