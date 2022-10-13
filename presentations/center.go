package presentations

import (
	"bytes"
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
	"unicode/utf8"
)

func NCenter(width int, s string, symbol string) *bytes.Buffer {
	// NCenter centers the string to the column width.
	const half = 2
	var b bytes.Buffer
	n := (width - utf8.RuneCountInString(s)) / half
	if n < 1 {
		fmt.Fprintf(&b, s)
		return &b
	}
	fmt.Fprintf(&b, "%s%s%s", strings.Repeat(symbol, int(n)), s, strings.Repeat(symbol, int(n)))
	return &b
}

// Center the string to the width of the terminal.
// When the width is unknown, the string is left-aligned.
func Center(s string, symbol string) *bytes.Buffer {
	fd := int(os.Stdin.Fd())
	w, _, err := term.GetSize(fd)
	if err != nil {
		return NCenter(0, s, symbol)
	}
	return NCenter(w, s, symbol)
}
