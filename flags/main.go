package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printHelp() {
	helpText := `--insert
  -i
	 This flag inserts the string into the string passed as argument.
--order
  -o
	 This flag will behave like a boolean, if it is called it will order the argument.
`
	printStr(helpText)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	insertStr := ""
	orderFlag := false
	mainStr := ""

	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			printHelp()
			return
		} else if arg == "-o" || arg == "--order" {
			orderFlag = true
		} else if len(arg) >= 8 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else {
			mainStr = arg
		}
	}

	finalStr := mainStr + insertStr

	if orderFlag {
		runes := []rune(finalStr)
		for i := 0; i < len(runes); i++ {
			for j := i + 1; j < len(runes); j++ {
				if runes[i] > runes[j] {
					runes[i], runes[j] = runes[j], runes[i]
				}
			}
		}
		finalStr = string(runes)
	}

	printStr(finalStr)
	z01.PrintRune('\n')
}
