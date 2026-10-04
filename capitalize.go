package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i, r := range runes {
		isLetter := (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
		isDigit := r >= '0' && r <= '9'
		if !isLetter && !isDigit {
			newWord = true
			continue
		}

		if isLetter {
			if newWord {
				if r >= 'a' && r <= 'z' {
					runes[i] = r - ('a' - 'A')
				}
			} else if r >= 'A' && r <= 'Z' {
				runes[i] = r + ('a' - 'A')
			}
		}
		newWord = false
	}

	return string(runes)
}
