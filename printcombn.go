package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	var current []rune
	isFirst := true

	var backtrack func(start int)
	backtrack = func(start int) {
		// Base case: when we have selected n digits
		if len(current) == n {
			if !isFirst {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			isFirst = false

			for _, r := range current {
				z01.PrintRune(r)
			}
			return
		}

		// Choose the next digit strictly greater than the previous
		for i := start; i <= 9; i++ {
			current = append(current, rune('0'+i))
			backtrack(i + 1)
			current = current[:len(current)-1] // Backtrack
		}
	}

	backtrack(0)
	z01.PrintRune('\n')
}
