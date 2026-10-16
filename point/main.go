package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}
	setPoint(points)
	for i, n := range []int{points.x, points.y} {
		if i == 0 {
			printStr("x = ")
		} else {
			printStr(", y = ")
		}
		for _, digit := range []int{n / 10, n % 10} {
			for value, r := range "0123456789" {
				if value == digit {
					z01.PrintRune(r)
				}
			}
		}
	}
	printStr("\n")
}
