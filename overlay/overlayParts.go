package overlay

import (
	"fmt"
	"hangman-classic/presentations"
)

func PrintTop() func() {
	return func() {
		topLogo := []string{
			"______  __                                               ",
			"___  / / /_____ ______________ _______ _________ _______ ",
			"__  /_/ /_  __ `/_  __ \\_  __ `/_  __ `__ \\  __ `/_  __ \\",
			"_  __  / / /_/ /_  / / /  /_/ /_  / / / / / /_/ /_  / / /",
			"/_/ /_/  \\__,_/ /_/ /_/_\\__, / /_/ /_/ /_/\\__,_/ /_/ /_/ ",
			"                       /____/                            "}

		for i := 0; i < len(topLogo); i++ {
			for j := 0; j <= 50; j++ {
				topLogo[i] = " " + topLogo[i] + " "
			}
		}

		for i := 0; i < len(topLogo); i++ {
			fmt.Printf("%s\n", presentations.Center(topLogo[i], "\u0020")) // "\u0020" = space unicode
		}
	}
}

func PrintBottom() func() {
	return func() {
		bottomLogo := "good luck"
		fmt.Printf("%s", presentations.Center(bottomLogo, "\u003D")) // "\u003D" = equal unicode
	}
}
