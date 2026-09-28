package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}

	sub := RecursiveFactorial(nb - 1)
	if sub == 0 {
		return 0
	}

	result := nb * sub

	if result/nb != sub {
		return 0
	}

	return result
}
