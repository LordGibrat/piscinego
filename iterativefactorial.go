package piscine

func IterativeFactorial(nb int) int {
	for i := 1; i <= nb; i++ {
		Ans = i * Ans
		return Ans
	}
}
