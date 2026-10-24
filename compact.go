package piscine

func Compact(ptr *[]string) int {
	if ptr == nil {
		return 0
	}
	count := 0
	for _, value := range *ptr {
		if value != "" {
			(*ptr)[count] = value
			count++
		}
	}
	*ptr = (*ptr)[:count]
	return count
}
