package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	upper := false
	if args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		n := 0
		valid := true

		if len(arg) == 0 {
			valid = false
		} else {
			for _, ch := range arg {
				if ch < '0' || ch > '9' {
					valid = false
					break
				}
				n = n*10 + int(ch-'0')
			}
		}

		if valid && n >= 1 && n <= 26 {
			if upper {
				z01.PrintRune(rune('A' + n - 1))
			} else {
				z01.PrintRune(rune('a' + n - 1))
			}
		} else {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}
