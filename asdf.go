// Package piscine contains the exercise helper.
package piscine

import "fmt"

func main() {
	var name string
	var age int

	fmt.Scan(&name, &age)
	fmt.Println("My name is", name, "and I am", age, "years old.")
}
