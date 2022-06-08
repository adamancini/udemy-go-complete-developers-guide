package main

import (
	"fmt"
	"io"
	"os"
)

// read teh contents of a file on the local drive and print the contents to the terminal
// file to open should be provided as an arugment to the program when it is executed at terminal
// example: go run main.go myfile.txt should open up myfile.txt
// to read in the arguments, you can reference os.Args which is a []string
// open a file: check out Open function in os package

// type fileReader struct{}

// func (fileReader) Read(bs []byte) (int, error) {

// }

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
