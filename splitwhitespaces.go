package piscine

func SplitWhiteSpaces(s string) []string {
	result := make([]string, 0)
	start := -1
	for i, r := range s {
		separator := r == ' ' || r == '\t' || r == '\n'
		if separator {
			if start >= 0 {
				result = append(result, s[start:i])
				start = -1
			}
		} else if start < 0 {
			start = i
		}
	}
	if start >= 0 {
		result = append(result, s[start:])
	}
	return result
}
