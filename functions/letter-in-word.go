package functions

func LetterInWorld(word string, letter string) []int {
	// return the indexes of the letter in the word
	var indexes []int
	for i, l := range word {
		if string(l) == letter {
			indexes = append(indexes, i)
		}
	}
	return indexes
}