package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && f(a[j-1], a[j]) > 0; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
