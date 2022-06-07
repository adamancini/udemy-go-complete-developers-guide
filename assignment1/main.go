package main

import "fmt"

func main() {
	slice := make([]int, 11)
	for i := 0; i <= 10; i++ {
		slice[i] = i
	}

	for _, v := range slice {
		if v%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
