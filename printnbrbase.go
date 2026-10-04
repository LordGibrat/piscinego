package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	digits := []rune(base)
	if len(digits) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < len(digits); i++ {
		if digits[i] == '+' || digits[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < len(digits); j++ {
			if digits[i] == digits[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	value := int64(nbr)
	if value < 0 {
		z01.PrintRune('-')
		value = -value
	}
	if value == 0 {
		z01.PrintRune(digits[0])
		return
	}

	result := make([]rune, 0)
	baseSize := int64(len(digits))
	for value > 0 {
		result = append(result, digits[value%baseSize])
		value /= baseSize
	}
	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])
	}
}
