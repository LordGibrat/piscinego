//go:build ignore

package main

import "fmt"

func main() {
	sum := 0
	for count := 0; count <= 10; count++ {
		sum += count
	}
	fmt.Println(sum)
}
