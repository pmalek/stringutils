package stringutils

import (
	"testing"
)

type stripCharsTestData struct {
	str      string
	c        string
	expected string
}

func Test_Stripchars(t *testing.T) {
	var stripCharsTests = []stripCharsTestData{
		{"qwerty.asdfgh.", ".", "qwertyasdfgh"},
		{"qwerty.asdfgh.", ",", "qwerty.asdfgh."},
		{"qwerty.asdfgh.", " ", "qwerty.asdfgh."},
		{"qwerty.asdfgh.", "@", "qwerty.asdfgh."},
		{"qwerty.asdfgh.", "a", "qwerty.sdfgh."},
		{"qwerty.asdfgh.", "b", "qwerty.asdfgh."},
		{"abc", ".qwe", "abc"},
		{"123456", "123456", ""},
		{"1234567", "123456", "7"},
		{"1234567", "123456.", "7"},
		{"1234567", "123456abcdefghijklmno.", "7"},
		{"1234567", "123456abcdefghijklmno.123456abcdefghijklmno.123456abcdefghijklmno.123456abcdefghijklmno.", "7"},
	}

	for _, tt := range stripCharsTests {
		actual := Stripchars(tt.str, tt.c)
		if actual != tt.expected {
			t.Errorf("Stripchars(%q, %q): Expected: %q, actual %q", tt.str, tt.c, tt.expected, actual)
		}
	}
}

type padTestData struct {
	s        string
	padStr   string
	pLen     int
	expected string
}

func Test_LeftPad(t *testing.T) {
	var leftPadTests = []padTestData{
		{"aa", "0", 3, "000aa"},
		{"abc", "g", 3, "gggabc"},
		{"1", "0", 0, "1"},
		{"1", "0", 50, "000000000000000000000000000000000000000000000000001"},
		{"TEST", "0", 5, "00000TEST"},
		{"TEST", "xx", 3, "xxxxxxTEST"},
		{"TEST", "xx", 25, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxTEST"},
	}

	for _, tt := range leftPadTests {
		actual := LeftPad(tt.s, tt.padStr, tt.pLen)
		if actual != tt.expected {
			t.Errorf("LeftPad(%q, %q): Expected: %q, actual %q", tt.s, tt.padStr, tt.expected, actual)
		}
	}
}

func Test_RightPad(t *testing.T) {
	var rightPadTests = []padTestData{
		{"aa", "0", 3, "aa000"},
		{"abc", "g", 3, "abcggg"},
		{"1", "0", 0, "1"},
		{"TEST", "0", 5, "TEST00000"},
		{"TEST", "xx", 3, "TESTxxxxxx"},
		{"TEST", "xx", 20, "TESTxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"},
	}

	for _, tt := range rightPadTests {
		actual := RightPad(tt.s, tt.padStr, tt.pLen)
		if actual != tt.expected {
			t.Errorf("RightPad(%q, %q): Expected: %q, actual %q", tt.s, tt.padStr, tt.expected, actual)
		}
	}
}

func Test_LeftPad2Len(t *testing.T) {
	var leftPad2LenTests = []padTestData{
		{"aa", "0", 3, "0aa"},
		{"abc", "g", 5, "ggabc"},
		{"1", "0", 1, "1"},
		{"1", "0", 0, ""},
		{"TEST", "0", 5, "0TEST"},
		{"TEST", "0", 50, "0000000000000000000000000000000000000000000000TEST"},
	}

	for _, tt := range leftPad2LenTests {
		actual := LeftPad2Len(tt.s, tt.padStr, tt.pLen)
		if actual != tt.expected {
			t.Errorf("LeftPad2Len(%q, %q): Expected: %q, actual %q", tt.s, tt.padStr, tt.expected, actual)
		}
	}
}

func Test_RightPad2Len(t *testing.T) {
	var rightPad2LenTests = []padTestData{
		{"aa", "0", 3, "aa0"},
		{"abc", "g", 5, "abcgg"},
		{"1", "0", 1, "1"},
		{"1", "0", 0, ""},
		{"TEST", "0", 5, "TEST0"},
		{"TEST", "0", 50, "TEST0000000000000000000000000000000000000000000000"},
	}

	for _, tt := range rightPad2LenTests {
		actual := RightPad2Len(tt.s, tt.padStr, tt.pLen)
		if actual != tt.expected {
			t.Errorf("RightPad2Len(%q, %q): Expected: %q, actual %q", tt.s, tt.padStr, tt.expected, actual)
		}
	}
}
