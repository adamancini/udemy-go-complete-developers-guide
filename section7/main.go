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

	c := make(chan string)

	for _, link := range links {
		// checkLink(link)  // move to `go checkLInk()`
		go checkLink(link, c) // spawn a new goroutine when checkLink blocks on http.Get()
	}

	// fmt.Println(<-c)
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // http.Get() blocks - we want to fix this with goroutines
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep, it's up"
}
