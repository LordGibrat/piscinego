package piscine

func LoafOfBread(str string) string {
	letters := []rune(str)
	if len(letters) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	for i := 0; i < len(letters); {
		word := ""
		count := 0
		for i < len(letters) && count < 5 {
			if letters[i] != ' ' {
				word += string(letters[i])
				count++
			}
			i++
		}
		if word != "" {
			if result != "" {
				result += " "
			}
			result += word
		}
		i++
	}
	return result + "\n"
}
