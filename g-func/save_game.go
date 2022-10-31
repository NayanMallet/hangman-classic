package g_func

import (
	"encoding/json"
	"fmt"
	"os"
)

type NewSave struct {
	Word             string
	WordRune         []rune
	Attempts         int
	LettersSuggested []string
	LetterFile       string
}

func SaveGame(save *NewSave) {
	// Function who create a save of the game in a file "save.txt"
	b, err := json.Marshal(save)
	if err != nil {
		fmt.Println("error:", err)
	}
	file, err2 := os.Create("save.txt")
	defer file.Close() // Close the file at the end of the program
	if err2 != nil {
		fmt.Println("error:", err2)
	}

	_, err3 := file.WriteString(string(b)) // Write in the file
	if err3 != nil {
		fmt.Println("error:", err3)
	}
}
