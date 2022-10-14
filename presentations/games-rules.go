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
