package g_func

func ContainsString(s string, car string) bool {
	// function to check if a character is present in a string
	for _, c := range s {
		if string(c) == car {
			return true
		}
	}
	return false
}

func ContainsTable(s []string, car string) bool {
	// function to check if a character is present in a table of  string
	for _, c := range s {
		if c == car {
			return true
		}
	}
	return false
}
