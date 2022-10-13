package presentations

import "time"

func Sleep(i int) {
	// This function will be used to sleep the program for a given time
	time.Sleep(time.Duration(i) * time.Second)
}
