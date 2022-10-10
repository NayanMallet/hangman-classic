package functions

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Game() {
	if os.Args[1] == "--startWith" && os.Args[2] == "save.txt" {

	}
	word := strings.ToUpper(RandomWord(os.Args[1]))
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	var hiddenLetters []int   // stocking the indices of the hidden letters
	var revealedLetters []int // stocking the indices of the revealed letters
	wordRune := []rune(word)

	for i := 0; i < len(word)/2-1; i++ {
		randTemp := randSource.Intn(len(wordRune))
		if wordRune[randTemp] != 0 {
			revealedLetters = append(revealedLetters, randTemp)
			wordRune[randTemp] = 0
		} else {
			i--
		}
	}

	for j := 0; j < len(wordRune); j++ {
		if wordRune[j] != 0 {
			hiddenLetters = append(hiddenLetters, j)
		}
	}

	for _, i := range revealedLetters {
		wordRune[i] = rune(word[i])
	}

	for _, i := range hiddenLetters {
		wordRune[i] = '_'
	}

	attempts := 10
	fmt.Printf("Good Luck, you have 10 attempts.\n%s\n", string(wordRune))
	for attempts > 0 {
		if string(wordRune) == word {
			fmt.Printf("%s\n", "Congrats !")
			return
		}
		var letter string

		_, err := fmt.Scan(&letter)
		if err != nil {
			fmt.Println(err)
		}

		letter = strings.ToUpper(letter)
		if len(letter) > 1 {
			if letter == word {
				fmt.Printf("%s\n", "Congrats !")
				return
			} else {
				attempts--
				fmt.Printf("Wrong ! You have %d attempts left.\n%s\n", attempts, string(wordRune))
				PrintMan(attempts)
			}
		} else {
			if Contains(word, letter) {
				indexes := LetterInWorld(word, letter)
				for _, i := range indexes {
					wordRune[i] = rune(word[i])
				}
				fmt.Printf("Choose: %s\n%s\n", letter, string(wordRune))
			} else {
				attempts--
				fmt.Printf("Choose: %s\nNot present in the word, %v attempts remaining\n", letter, attempts)
				PrintMan(attempts)
			}
		}
	}
	fmt.Printf("%s\n", "You Loose !")
}
