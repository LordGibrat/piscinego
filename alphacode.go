package piscine

func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		// Проверяем, является ли символ строчной или заглавной буквой латинского алфавита
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			count++
		}
	}
	return count
}
