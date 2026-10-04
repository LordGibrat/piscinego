package piscine

func TrimAtoi(s string) int {
	sign := 1
	value := 0
	foundDigit := false

	for _, r := range s {
		if r == '-' && !foundDigit {
			sign = -1
			continue
		}
		if r == '+' && !foundDigit {
			sign = 1
			continue
		}
		if r < '0' || r > '9' {
			continue
		}

		foundDigit = true
		value = value*10 + int(r-'0')
	}

	if !foundDigit {
		return 0
	}
	return sign * value
}
