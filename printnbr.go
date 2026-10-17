// Package piscine contains the exercise helper.
package piscine

import "fmt"

func PrintNbrMain() {
	sum := 0
	for count := 0; count <= 10; count++ {
		sum += count
	}
	fmt.Println(sum)
}
