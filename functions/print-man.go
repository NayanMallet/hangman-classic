package functions

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func PrintMan(attempts int) {
	data, err := ioutil.ReadFile("man-positions/" + strconv.Itoa(attempts) + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", string(data))
}
