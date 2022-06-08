package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://ifconfig.me",
		"http://wikipedia.org",
		"http://golang.com",
		"http://google.com",
	}

	for _, link := range links {
		// checkLink(link)
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link) // http.Get() blocks - we want to fix this with goroutines
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}
