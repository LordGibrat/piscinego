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

	for _, name := range os.Args[1:] {
		file, err := os.Open(name)
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			os.Exit(1)
		}
		if _, err = io.Copy(os.Stdout, file); err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
			file.Close()
			os.Exit(1)
		}
		file.Close()
	}
}
