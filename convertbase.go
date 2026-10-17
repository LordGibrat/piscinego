package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	from := []rune(baseFrom)
	to := []rune(baseTo)
	value := 0
	for _, digit := range []rune(nbr) {
		position := 0
		for i, baseDigit := range from {
			if digit == baseDigit {
				position = i
				break
			}
		}
		value = value*len(from) + position
	}
	if value == 0 {
		return string(to[0])
	}

	result := make([]rune, 0)
	for value > 0 {
		result = append(result, to[value%len(to)])
		value /= len(to)
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
