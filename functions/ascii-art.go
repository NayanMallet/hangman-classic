package functions

import (
	"bufio"
	"fmt"
	"os"
)

func AsciiArt(s string, file string) {
	table := []rune(s)
	if file, err := os.Open(file); err != nil {
		fmt.Printf("Error: %s", err)
		return
	} else {
		defer file.Close()
		var lines []string
		for scanner := bufio.NewScanner(file); scanner.Scan(); {
			lines = append(lines, scanner.Text())
		}
		var place []int
		for _, car := range table {
			if car == 32 {
				place = append(place, 0)
			} else {
				place = append(place, (int(car)-32)*9)
			}
		}
		for i := 0; i < 10; i++ {
			lineToPrint := ""
			for j, places := range place {
				lineToPrint += lines[places]
				place[j] += 1
			}
			fmt.Printf("%s\n", lineToPrint)
		}

	}
}

// 9 -> length car
