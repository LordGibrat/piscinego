package piscine

func IterativeFactorial(nb int) int {
	Ans := 1
	if nb < 0 {
		return 0
	}
	if nb != int(nb) {
		return 0
	}
	for i := 1; i <= nb; i++ {
		Ans = i * Ans
	}
	return Ans
}
