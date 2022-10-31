package g_func

import (
	"math/rand"
	"strings"
	"time"
)

func NewGamePrep(args []string) (word string, wordRune []rune) {
	// Preparing the game with default settings

	word = strings.ToUpper(RandomWord(args[0]))
	// word = random word
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	var hiddenLetters, revealedLetters []int
	// stocking the indices of the hidden/revealed letters

	wordRune = []rune(word)
	// splitting the word into runes

	for i := 0; i < len(word)/2-1; i++ {
		// revealing half of the letters
		randTemp := randSource.Intn(len(wordRune))
		if wordRune[randTemp] != 0 {
			revealedLetters = append(revealedLetters, randTemp)
			wordRune[randTemp] = 0
		} else {
			i--
		}
	}
	for j := 0; j < len(wordRune); j++ {
		// stocking the indices of the hidden letters
		if wordRune[j] != 0 {
			hiddenLetters = append(hiddenLetters, j)
		}
	}
	for _, i := range revealedLetters {
		// revealing the letters
		wordRune[i] = rune(word[i])
	}
	for _, i := range hiddenLetters {
		// hiding the letters
		wordRune[i] = '_'
	}
	return word, wordRune
}
