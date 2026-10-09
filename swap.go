package piscine

func Swap(a *int, b *int) {
	buf := *a
	*a = *b
	*b = buf
}
