package main

import (
	"fmt"
	"net/http"
	"time"
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
		// fmt.Println(<-c)
	}

	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)
	// }

	// infinite loop
	// for {
	// 	go checkUrl(<-c, c) // go routine being called repeatedly
	// }

	// Alternate Syntax
	for l := range c { // loop through everytime a channel emits a value
		// time.Sleep(5 * time.Second) // pause for 5 seconds // main routine is getting paused here
		// go checkUrl(l, c)
		// go func() { //function literal
		// 	time.Sleep(5 * time.Second)
		// 	checkUrl(l, c)
		// }()

		go func(link string) { // recommended; 'l' string is copied in memory, Go routine has access to that copy instead of pointing to the original value of l
			time.Sleep(5 * time.Second)
			checkUrl(link, c)
		}(l)
	}
}

func checkUrl(url string, c chan string) {
	// _, err := http.Get(url) // blocking call; add Go routines
	// time.Sleep(5 * time.Second) // go routine is waiting here; checkUrl should typically jusrt execute, it doesn't make sense for it to wait. We need something between main and checkUrl to wait
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
		// c <- "Might be down"
		c <- url
		return
	}
	fmt.Println(url, "is up")
	// c <- "Yes, it's up!"
	c <- url
}
