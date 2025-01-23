package gstr

import "strings"

// Repeat returns a new string consisting of multiplier copies of the string input.
func Repeat(input string, multiplier int) string {
	return strings.Repeat(input, multiplier)
}
