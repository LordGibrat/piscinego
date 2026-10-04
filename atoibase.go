package piscine

func AtoiBase(s string, base string) int {
	digits := []rune(base)
	if len(digits) < 2 {
		return 0
	}
	for i, digit := range digits {
		if digit == '+' || digit == '-' {
			return 0
		}
		for _, other := range digits[i+1:] {
			if digit == other {
				return 0
			}
		}
	}

	value := 0
	for _, digit := range []rune(s) {
		position := -1
		for i, baseDigit := range digits {
			if digit == baseDigit {
				position = i
				break
			}
		}
		if position < 0 {
			return 0
		}
		value = value*len(digits) + position
	}
	return value
}
