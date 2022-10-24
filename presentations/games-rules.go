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
	fmt.Printf(
		"%s\n%s\n",
		Center("If you want to start a new game, type 'NEW'.", "\u0020"),
		// type 'NEW' to start a new game
		Center("If you want to load last save, type 'LOAD'.", "\u0020"))
	// type 'LOAD' to Load last game
}

func GameOptionChoice() {
	fmt.Printf(
		"%s\n%s\n",
		Center("For load custom Words list enter file name (ex: yourFile.txt)\n", "\u0020"),
		// type 'yourFile.txt' to load a custom word list
		Center("Else tap 'N'\n", "\u0020"))
}

func GameAsciiOptionChoice() {
	fmt.Printf(
		"%s\n%s\n",
		Center("For load custom Ascii-Art file enter file name (ex: yourFile.txt)\n", "\u0020"),
		// type 'yourFile.txt' to load a custom Ascii-Art Newfile
		Center("Else tap 'N'\n", "\u0020"))
}
