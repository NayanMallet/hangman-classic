package presentations

import (
	"bufio"
	"fmt"
	"os"
)

func PrintMan(attempts int) {
	if file, err := os.Open("hangman.txt"); err != nil {
		fmt.Printf("Error: %s", err)
		return
	} else {
		defer file.Close()
		var lines []string
		for scanner := bufio.NewScanner(file); scanner.Scan(); {
			lines = append(lines, scanner.Text())
		}
		if attempts == 9 {
			for i := 0; i < 8; i++ {
				fmt.Println(lines[i])
			}
		} else if attempts == 8 {
			for i := 8; i < 16; i++ {
				fmt.Println(lines[i])
			}
		} else if attempts == 7 {
			for i := 16; i < 24; i++ {
				fmt.Println(lines[i])
			}
		} else if attempts == 6 {
			for i := 24; i < 32; i++ {
				fmt.Println(lines[i])
			}
		} else if attempts == 5 {
			for i := 32; i < 40; i++ {
				fmt.Println(lines[i])
			}
		} else if attempts == 4 {
			for i := 40; i < 48; i++ {
				fmt.Println(lines[i])
			}
		} else if attempts == 3 {
			for i := 48; i < 56; i++ {
				fmt.Println(lines[i])
			}
		} else if attempts == 2 {
			for i := 56; i < 64; i++ {
				fmt.Println(lines[i])
			}
		} else if attempts == 1 {
			for i := 64; i < 72; i++ {
				fmt.Println(lines[i])
			}
		} else if attempts == 0 {
			for i := 72; i < 80; i++ {
				fmt.Println(lines[i])
			}
		}
	}
}
