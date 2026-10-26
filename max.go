package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	max := a[0]
	for _, value := range a[1:] {
		if value > max {
			max = value
		}
	}
	return max
}
