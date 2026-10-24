package piscine

func JumpOver(str string) string {
	result := ""
	for i, r := range []rune(str) {
		if i%3 == 2 {
			result += string(r)
		}
	}
	return result + "\n"
}
