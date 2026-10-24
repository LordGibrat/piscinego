package piscine

func Abort(a, b, c, d, e int) int {
	values := [5]int{a, b, c, d, e}
	for i := 1; i < len(values); i++ {
		for j := i; j > 0 && values[j] < values[j-1]; j-- {
			values[j], values[j-1] = values[j-1], values[j]
		}
	}
	return values[2]
}
