package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
		// fmt.Println(<-c)
	}

	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-c, c)
	// }
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Function Literal
	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "might be down!")
		c <- l
		return
	}

	fmt.Println(l, "is up!")
	c <- l
}
