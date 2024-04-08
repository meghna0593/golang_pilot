package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Create a channel
	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c) // go routine
		fmt.Println(<-c)
	}

	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)
	// }
}

func checkUrl(url string, c chan string) {
	// _, err := http.Get(url) // blocking call; add Go routines
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
		c <- "Might be down"
		return
	}
	fmt.Println(url, "is up")
	c <- "Yes, it's up!"
}
