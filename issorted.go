package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending, descending := true, true
	for i := 1; i < len(a); i++ {
		comparison := f(a[i-1], a[i])
		if comparison > 0 {
			ascending = false
		}
		if comparison < 0 {
			descending = false
		}
	}
	return ascending || descending
}
