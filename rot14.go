package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = 'a' + (r-'a'+14)%26
		} else if r >= 'A' && r <= 'Z' {
			runes[i] = 'A' + (r-'A'+14)%26
		}
	}
	return string(runes)
}
