package g_func

import "fmt"

func GetInput() (letter string) {
	// function to retrieve the input of the user
	_, err := fmt.Scan(&letter)
	if err != nil {
		fmt.Println(err)
	}
	return letter
}
