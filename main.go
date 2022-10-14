package main

import (
	"hangman-classic/overlay"
)

func main() {
	overlay.PrintTop()
	overlay.PrintBottom()
	//overlay.Wrapper(overlay.PrintTop, game.Progress(), overlay.PrintBottom)()
}
