package main

import (
	"fmt"
	"os/exec"
)

func ziz() {
	fmt.Println("zizi")
}

func main() {
	//overlay.Wrapper(overlay.PrintTop, presentations.GamesRules(), overlay.PrintBottom)()
	cmd := exec.Command("go", "run", "hangman-classic/game", "words.txt")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}
