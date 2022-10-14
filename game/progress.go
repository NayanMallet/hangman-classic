package game

import "hangman-classic/presentations"

func Progress() func() {
	// This function will be used to print the progress of the game.
	return func() {
		presentations.GamesRules()
		presentations.Sleep(1)
		Game("words.txt")
	}
}
