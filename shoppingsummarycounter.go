package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	start := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			summary[str[start:i]]++
			start = i + 1
		}
	}
	summary[str[start:]]++
	return summary
}
