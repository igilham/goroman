package roman

import (
	"bytes"
)

var lookup = []int {
	1000, 900, 500,  400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
}

var numerals = []string {
	"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
}

// Returns the Roman numeral representation of the given integer.
// Returns the empty string for numbers out of the range of 0 to 3999.
// Returns 'N' if the number is '0'.
func NToRoman(number int) string {
	if number < 0 || number > 3999 {
		return ""
	} else if number == 0 {
		return "N"
	}
	var buf bytes.Buffer
	for i, v := range lookup {
		for number >= v {
			number -= v
			buf.WriteString(numerals[i])
		}
	}
	return buf.String()
}
