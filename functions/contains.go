package functions

func Contains(s string, car string) bool {
	// function to check if an character is present in a string
	for _, c := range s {
		if string(c) == car {
			return true
		}
	}
	return false
}
