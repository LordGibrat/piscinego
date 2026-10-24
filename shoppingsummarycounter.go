package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""
	for _, r := range str {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '\v' || r == '\f' {
			if word != "" {
				summary[word]++
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		summary[word]++
	}
	return summary
}
