package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Game() {
	if !(os.Args[1] == "--startWith" && os.Args[2] == "save.txt") {

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
		var lettersSuggested []string

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
				if letter == "STOP" {
					SaveGame(&NewSave{Word: word, WordRune: wordRune, Attempts: attempts, LettersSuggested: lettersSuggested})
					fmt.Printf("Choose: %s\n\nGame Saved in save.txt.\n", letter)
					return
				}
				if letter == word {
					fmt.Printf("%s\n", "Congrats !")
					return
				} else {
					attempts -= 2
					fmt.Printf("Wrong ! You have %d attempts left.\n%s\n", attempts, string(wordRune))
					PrintMan(attempts)
				}
			} else {
				if ContainsTable(lettersSuggested, letter) {
					fmt.Printf("You already suggested this letter !\n%s\n", string(wordRune))
				} else {
					if ContainsString(word, letter) {
						lettersSuggested = append(lettersSuggested, letter)
						indexes := LetterInWorld(word, letter)
						for _, i := range indexes {
							wordRune[i] = rune(word[i])
						}
						fmt.Printf("Choose: %s\n%s\n", letter, string(wordRune))
					} else {
						attempts--
						fmt.Printf("Choose: %s\nNot present in the Word, %v attempts remaining\n", letter, attempts)
						PrintMan(attempts)
					}
				}
			}
		}
		fmt.Printf("You Loose !\nThe word to find was %s", word)
	} else {
		data, err := ioutil.ReadFile("save.txt")
		if err != nil {
			fmt.Println(err)
		}
		var GameStats *NewSave
		if err2 := json.Unmarshal(data, &GameStats); err != nil {
			fmt.Println(err2)
		}
		word, wordRune, attempts, lettersSuggested := GameStats.Word, GameStats.WordRune, GameStats.Attempts, GameStats.LettersSuggested

		fmt.Printf("Welcome Back, you have %v attempts remaining.\n%s\n", attempts, string(wordRune))
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
				if letter == "STOP" {
					SaveGame(&NewSave{Word: word, WordRune: wordRune, Attempts: attempts, LettersSuggested: lettersSuggested})
					fmt.Printf("Choose: %s\n\nGame Saved in save.txt.\n", letter)
					return
				}
				if letter == word {
					fmt.Printf("%s\n", "Congrats !")
					return
				} else {
					attempts--
					fmt.Printf("Wrong ! You have %d attempts left.\n%s\n", attempts, string(wordRune))
					PrintMan(attempts)
				}
			} else {
				if ContainsTable(lettersSuggested, letter) {
					fmt.Printf("You already suggested this letter !\n%s\n", string(wordRune))
				} else {
					if ContainsString(word, letter) {
						lettersSuggested = append(lettersSuggested, letter)
						indexes := LetterInWorld(word, letter)
						for _, i := range indexes {
							wordRune[i] = rune(word[i])
						}
						fmt.Printf("Choose: %s\n%s\n", letter, string(wordRune))
					} else {
						attempts--
						fmt.Printf("Choose: %s\nNot present in the Word, %v attempts remaining\n", letter, attempts)
						PrintMan(attempts)
					}
				}
			}
		}
		fmt.Printf("You Loose !\nThe word to find was %s", word)
	}
}
