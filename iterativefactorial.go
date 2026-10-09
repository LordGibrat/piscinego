package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}

	ans := 1
	for i := 1; i <= nb; i++ {
		next := ans * i

		// Переполнение болғанын тексеру
		if next/i != ans {
			return 0
		}

		ans = next
	}

	return ans
}
