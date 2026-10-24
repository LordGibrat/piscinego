package piscine

func DescendAppendRange(max, min int) []int {
	var result []int
	for value := max; value > min; value-- {
		result = append(result, value)
	}
	return result
}
