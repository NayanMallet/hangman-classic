package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func randomWord() string {
	// Program will randomly choose a word in the file
	file, err := os.Open("words.txt")

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
	// return word of wordsTable at random indice
}

func main() {
	fmt.Println(randomWord())
}
