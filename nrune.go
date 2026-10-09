package piscine

func NRune(s string, n int) rune {
	if n <= 0 || n > len([]rune(s)) {
		return 0
	} else {
		return []rune(s)[n-1]
	}
}
