package piscine

func IterativeFactorial(nb int) int {
	Ans := 1
	for i := 1; i <= nb; i++ {
		Ans = i * Ans
	}
	return Ans
}
