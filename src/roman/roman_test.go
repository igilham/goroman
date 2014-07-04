package roman

import (
	"testing"
)

type romanTest struct {
	input int
	expected string 
}

var romanCases = []romanTest{
	romanTest{-1, ""}, // bad input
	romanTest{4000, ""}, // bad input
	romanTest{0, "N"},
	romanTest{1, "I"},
	romanTest{2, "II"},
	romanTest{3, "III"},
	romanTest{4, "IV"},
	romanTest{5, "V"},
	romanTest{6, "VI"},
	romanTest{7, "VII"},
	romanTest{8, "VIII"},
	romanTest{9, "IX"},
	romanTest{10, "X"}, // fails on my first implementation after here
	romanTest{11, "XI"},
	romanTest{40, "XL"},
	romanTest{49, "XLIX"},
	romanTest{50, "L"},
	romanTest{64, "LXIV"},
	romanTest{89, "LXXXIX"},
	romanTest{90, "XC"},
	romanTest{100, "C"},
	romanTest{128, "CXXVIII"},
	romanTest{399, "CCCXCIX"},
	romanTest{400, "CD"},
	romanTest{490, "CDXC"},
	romanTest{500, "D"},
	romanTest{888, "DCCCLXXXVIII"},
	romanTest{900, "CM"},
	romanTest{999, "CMXCIX"},
	romanTest{1000, "M"},
	romanTest{1947, "MCMXLVII"},
	romanTest{1999, "MCMXCIX"},
	romanTest{2000, "MM"},
	romanTest{3999, "MMMCMXCIX"},
}

func TestRoman(t *testing.T) {
	for _, rt := range romanCases {
		actual := NToRoman(rt.input)
		if actual != rt.expected  {
			t.Errorf("Parse(%v): expected \"%v\" but was \"%v\"\n", 
				rt.input, rt.expected, actual)
		}
	}
}
