package enum

import "strings"

// Repeat takes one characater and repeats it 5 times.
func Repeat(character string, times int) string {
	if times <= 0 {
		times = 5
	}

	return strings.Repeat(character, times)
}
