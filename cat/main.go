package main

import (
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	failed := false
	for _, name := range os.Args[1:] {
		file, err := os.Open(name)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			failed = true
			continue
		}
		if _, err = io.Copy(os.Stdout, file); err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			failed = true
		}
		file.Close()
	}
	if failed {
		os.Exit(1)
	}
}
