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
	b, err := json.Marshal(save)
	if err != nil {
		fmt.Println("error:", err)
	}
	file, err2 := os.Create("save.txt")
	defer file.Close() // on ferme automatiquement à la fin de notre programme
	if err2 != nil {
		fmt.Println("error:", err2)
	}

	_, err3 := file.WriteString(string(b)) // écrire dans le fichier
	if err3 != nil {
		fmt.Println("error:", err3)
	}
}
