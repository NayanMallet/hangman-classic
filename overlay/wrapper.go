package overlay

func Wrapper(top func(), f func(), bottom func()) {
	// This is a wrapper function that will be used to wrap the function
	top()
	f()
	bottom()
}
