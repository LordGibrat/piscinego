package piscine

func RecursiveFactorial(nb int) int {
	// 0-ден кіші сандарға және 20-дан үлкен сандарға (overflow) 0 қайтарамыз
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}

	sub := RecursiveFactorial(nb - 1)
	result := nb * sub

	// Переполнение сақтану тексерісі
	if result/nb != sub {
		return 0
	}

	return result
}
