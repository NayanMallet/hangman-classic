package g_func

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func RandomWord(pathFile string) string {
	// Program will randomly choose a Word in the file
	pathFile = "words-files/" + pathFile
	file, err := os.Open(pathFile)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var wordsTable []string
	// table stocking all words separated

	for scanner.Scan() {
		wordsTable = append(wordsTable, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	return wordsTable[randSource.Intn(len(wordsTable))]
	// return Word of wordsTable at random indice
}
