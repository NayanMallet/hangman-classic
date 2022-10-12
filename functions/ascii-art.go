package functions

import (
	"bufio"
	"fmt"
	"os"
)

func AsciiArt(s string) {
	table := []rune(s)
	fmt.Println(table)
	if file, err := os.Open("standard.txt"); err != nil {
		fmt.Printf("Error: %s", err)
		return
	} else {
		defer file.Close()
		var lines []string
		for scanner := bufio.NewScanner(file); scanner.Scan(); {
			lines = append(lines, scanner.Text())
		}
		for _, rune := range table {
			if rune == 32 {
				for i := 0; i < 8; i++ {
					fmt.Println(lines[i])
				}
			} else {
				for i := (rune - 32) * 8; i < 8; i++ {
					fmt.Println(lines[i])
				}
			}
		}
	}
}
