package piscine

func Split(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	result := make([]string, 0)
	start := 0
	for i := 0; i+len(sep) <= len(s); {
		match := true
		for j := 0; j < len(sep); j++ {
			if s[i+j] != sep[j] {
				match = false
				break
			}
		}
		if match {
			result = append(result, s[start:i])
			i += len(sep)
			start = i
		} else {
			i++
		}
	}
	return append(result, s[start:])
}
