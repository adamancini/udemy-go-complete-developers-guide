package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: pass a file to open as an argument")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: could not open file for reading: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
