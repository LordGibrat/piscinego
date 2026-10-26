package piscine

func ReverseMenuIndex(menu []string) []string {
	reversed := make([]string, len(menu))
	for i := range menu {
		reversed[i] = menu[len(menu)-1-i]
	}
	return reversed
}
