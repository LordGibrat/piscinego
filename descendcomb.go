package piscine

import "github.com/01-edu/z01"

func descendPrintTwoDigits(n int) {
	z01.PrintRune(rune('0' + n/10))
	z01.PrintRune(rune('0' + n%10))
}

func DescendComb() {
	for a := 99; a >= 1; a-- {
		for b := a - 1; b >= 0; b-- {
			descendPrintTwoDigits(a)
			z01.PrintRune(' ')
			descendPrintTwoDigits(b)
			if a != 1 || b != 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
