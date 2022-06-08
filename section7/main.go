package main

import (
	"fmt"
	"net/http"
	"time"
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
	// for i := 0; i < len(links); i++ {
	for l := range c {
		// fmt.Println(<-c)
		// go checkLink(l,c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// time.Sleep(5 * time.Second)
	_, err := http.Get(link) // http.Get() blocks - we want to fix this with goroutines
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
