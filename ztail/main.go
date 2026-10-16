package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 || os.Args[1] != "-c" {
		os.Exit(1)
	}

	count := 0
	for _, digit := range os.Args[2] {
		count = count*10 + int(digit-'0')
	}

	failed := false
	printed := false
	for _, name := range os.Args[3:] {
		data, err := os.ReadFile(name)
		if err != nil {
			fmt.Printf("%v\n", err)
			failed = true
			continue
		}
		if len(os.Args) > 4 {
			if printed {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", name)
		}
		start := len(data) - count
		if start < 0 {
			start = 0
		}
		fmt.Printf("%s", data[start:])
		printed = true
	}
	if failed {
		os.Exit(1)
	}
}
