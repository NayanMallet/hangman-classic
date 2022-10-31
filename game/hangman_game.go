package game

import (
	"encoding/json"
	"fmt"
	"hangman-classic/g-func"
	"hangman-classic/presentations"
	"io/ioutil"
	"strings"
)

func Game(arg string) {
	args := strings.Split(arg, " ")
	// args = arguments

	switch len(args) {
	case 1:
		// New game with default settings
		word, wordRune := g_func.NewGamePrep(args)
		attempts := 10
		var lettersSuggested []string
		// lettersSuggested = letters suggested by the player

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
					g_func.SaveGame(&g_func.NewSave{Word: word, WordRune: wordRune, Attempts: attempts, LettersSuggested: lettersSuggested, LetterFile: ""})
					fmt.Printf("Choose: %s\n\nGame Saved in save.txt.\n", letter)
					return
				}
				if letter == word {
					fmt.Printf("%s\n", "Congrats !")
					return
				} else {
					attempts -= 2
					fmt.Printf("Wrong ! You have %d attempts left.\n%s\n", attempts, string(wordRune))
					presentations.PrintMan(attempts)
				}
			} else {
				if g_func.ContainsTable(lettersSuggested, letter) {
					fmt.Printf("You already suggested this letter !\n%s\n", string(wordRune))
				} else {
					if g_func.ContainsString(word, letter) {
						lettersSuggested = append(lettersSuggested, letter)
						indexes := g_func.LetterInWorld(word, letter)
						for _, i := range indexes {
							wordRune[i] = rune(word[i])
						}
						fmt.Printf("Choose: %s\n%s\n", letter, string(wordRune))
					} else {
						attempts--
						fmt.Printf("Choose: %s\nNot present in the Word, %v attempts remaining\n", letter, attempts)
						presentations.PrintMan(attempts)
					}
				}
			}
		}
		fmt.Printf("You Loose !\nThe word to find was %s", word)

	case 2:
		// Load the last game
		if args[0] == "--startWith" && args[1] == "save.txt" {
			data, err := ioutil.ReadFile("save.txt")
			// read the file
			if err != nil {
				fmt.Println(err)
			}
			var GameStats *g_func.NewSave
			if err2 := json.Unmarshal(data, &GameStats); err != nil {
				fmt.Println(err2)
			}
			word, wordRune, attempts, lettersSuggested, LetterFile := GameStats.Word, GameStats.WordRune, GameStats.Attempts, GameStats.LettersSuggested, GameStats.LetterFile
			// load all the stats of the game

			if LetterFile == "" {
				// Ascii art option is not activated
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
							g_func.SaveGame(&g_func.NewSave{Word: word, WordRune: wordRune, Attempts: attempts, LettersSuggested: lettersSuggested, LetterFile: LetterFile})
							fmt.Printf("Choose: %s\n\nGame Saved in save.txt.\n", letter)
							return
						}
						if letter == word {
							fmt.Printf("%s\n", "Congrats !")
							return
						} else {
							attempts--
							fmt.Printf("Wrong ! You have %d attempts left.\n%s\n", attempts, string(wordRune))
							presentations.PrintMan(attempts)
						}
					} else {
						if g_func.ContainsTable(lettersSuggested, letter) {
							fmt.Printf("You already suggested this letter !\n%s\n", string(wordRune))
						} else {
							if g_func.ContainsString(word, letter) {
								lettersSuggested = append(lettersSuggested, letter)
								indexes := g_func.LetterInWorld(word, letter)
								for _, i := range indexes {
									wordRune[i] = rune(word[i])
								}
								fmt.Printf("Choose: %s\n%s\n", letter, string(wordRune))
							} else {
								attempts--
								fmt.Printf("Choose: %s\nNot present in the Word, %v attempts remaining\n", letter, attempts)
								presentations.PrintMan(attempts)
							}
						}
					}
				}
				fmt.Printf("You Loose !\nThe word to find was %s", word)
			} else {
				// Ascii art option is activated
				fmt.Printf("Welcome Back, you have %v attempts remaining.\n", attempts)
				presentations.AsciiArt(string(wordRune), LetterFile)
				for attempts > 0 {
					if string(wordRune) == word {
						presentations.AsciiArt("Congrats !\n", LetterFile)
						return
					}
					var letter string
					_, err := fmt.Scan(&letter)
					if err != nil {
						fmt.Println(err)
					}
					letter = strings.ToUpper(letter)
					if len(letter) > 1 {
						// if the user type more than one letter
						if letter == "STOP" {
							g_func.SaveGame(&g_func.NewSave{Word: word, WordRune: wordRune, Attempts: attempts, LettersSuggested: lettersSuggested, LetterFile: LetterFile})
							fmt.Printf("Choose: %s\n\nGame Saved in save.txt.\n", letter)
							return
						}
						if letter == word {
							presentations.AsciiArt("Congrats !", LetterFile)
							return
						} else {
							attempts--
							fmt.Printf("Wrong ! You have %d attempts left.\n", attempts)
							presentations.AsciiArt(string(wordRune), LetterFile)
							presentations.PrintMan(attempts)
						}
					} else {
						// if the user type only one letter
						if g_func.ContainsTable(lettersSuggested, letter) {
							fmt.Printf("You already suggested this letter !\n")
							presentations.AsciiArt(string(wordRune), LetterFile)
						} else {
							if g_func.ContainsString(word, letter) {
								lettersSuggested = append(lettersSuggested, letter)
								indexes := g_func.LetterInWorld(word, letter)
								for _, i := range indexes {
									wordRune[i] = rune(word[i])
								}
								fmt.Printf("Choose: %s\n", letter)
								presentations.AsciiArt(string(wordRune), LetterFile)
							} else {
								attempts--
								fmt.Printf("Choose: %s\nNot present in the Word, %v attempts remaining\n", letter, attempts)
								presentations.PrintMan(attempts)
							}
						}
					}
				}
				fmt.Printf("You Loose !\nThe word to find was \n")
				presentations.AsciiArt(word, LetterFile)
			}
		} else {
			fmt.Println("Impossible...")
			break
		}

	case 3:
		// New game with Ascii art option
		if args[1] == "--letterFile" {
			word, wordRune := g_func.NewGamePrep(args)
			attempts := 10
			var lettersSuggested []string
			fmt.Printf("Good Luck, you have 10 attempts.\n")
			presentations.AsciiArt(string(wordRune), args[2])

			for attempts > 0 {
				if string(wordRune) == word {
					presentations.AsciiArt("Congrats !", args[2])
					return
				}
				var letter string
				_, err := fmt.Scan(&letter)
				if err != nil {
					fmt.Println(err)
				}
				letter = strings.ToUpper(letter)
				if len(letter) > 1 {
					// if the user type more than one letter
					if letter == "STOP" {
						g_func.SaveGame(&g_func.NewSave{Word: word, WordRune: wordRune, Attempts: attempts, LettersSuggested: lettersSuggested, LetterFile: args[2]})
						fmt.Printf("Choose: %s\n\nGame Saved in save.txt.\n", letter)
						return
					}
					if letter == word {
						presentations.AsciiArt("Congrats !", args[2])
						return
					} else {
						attempts -= 2
						fmt.Printf("Wrong ! You have %d attempts left.\n", attempts)
						presentations.AsciiArt(string(wordRune), args[2])
						presentations.PrintMan(attempts)
					}
				} else {
					// if the user type only one letter
					if g_func.ContainsTable(lettersSuggested, letter) {
						fmt.Printf("You already suggested this letter !\n")
						presentations.AsciiArt(string(wordRune), args[2])
					} else {
						if g_func.ContainsString(word, letter) {
							lettersSuggested = append(lettersSuggested, letter)
							indexes := g_func.LetterInWorld(word, letter)
							for _, i := range indexes {
								wordRune[i] = rune(word[i])
							}
							fmt.Printf("Choose: %s\n", letter)
							presentations.AsciiArt(string(wordRune), args[2])
						} else {
							attempts--
							fmt.Printf("Choose: %s\nNot present in the Word, %v attempts remaining\n", letter, attempts)
							presentations.PrintMan(attempts)
						}
					}
				}
			}
			fmt.Printf("You Loose !\nThe word to find was\n")
			presentations.AsciiArt(word, args[2])
		} else {
			fmt.Println("Impossible...")
			break
		}
	}
}
