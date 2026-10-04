package piscine 

func AlphaCount(s string) int {
	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			count++
		}
	}
	return count
}
