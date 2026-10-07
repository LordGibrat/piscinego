package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	var runes []rune
	for i, arg := range args {
		for _, r := range arg {
			runes = append(runes, r)
		}
		if i < len(args)-1 {
			runes = append(runes, ' ')
		}
	}

	left := 0
	right := len(runes) - 1

	for left < right {
		for left < right && !isVowel(runes[left]) {
			left++
		}
		for left < right && !isVowel(runes[right]) {
			right--
		}

		if left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}

	for _, r := range runes {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
