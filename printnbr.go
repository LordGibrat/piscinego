package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
	}

	printDigits(n)
}

func printDigits(n int) {
	if n == 0 {
		return
	}

	digit := n % 10
	if digit < 0 {
		digit = -digit
	}

	if n/10 != 0 {
		printDigits(n / 10)
	}

	z01.PrintRune(rune(digit + '0'))
}
