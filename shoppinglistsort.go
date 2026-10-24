package piscine

func ShoppingListSort(slice []string) []string {
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0 && len(slice[j]) < len(slice[j-1]); j-- {
			slice[j], slice[j-1] = slice[j-1], slice[j]
		}
	}
	return slice
}
