package piscine

func Unmatch(a []int) int {
	for _, value := range a {
		count := 0
		for _, candidate := range a {
			if candidate == value {
				count++
			}
		}
		if count%2 != 0 {
			return value
		}
	}
	return -1
}
