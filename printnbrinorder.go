package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var counts [10]int
	for n > 0 {
		counts[n%10]++
		n /= 10
	}

	for digit, count := range counts {
		for i := 0; i < count; i++ {
			z01.PrintRune(rune('0' + digit))
		}
	}
}
