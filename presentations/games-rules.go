package presentations

import (
	"fmt"
)

func GamesRules() func() {
	// print the rules of the hangman game
	return func() {
		rules := []string{
			"Rules are simple, you have to guess the word by suggesting letters.",
			"If you suggest a letter that is in the word, it will be revealed.",
			"If you suggest a letter that is not in the word, you will lose an attempt.",
			"If tou suggest the whole word, you will win if it is correct.",
			"If you lose all your attempts, you will lose the game."}
		for i := 0; i < len(rules); i++ {
			fmt.Printf("%s\n", Center(rules[i], "\u0020"))
			Sleep(1)
		}
	}
}
