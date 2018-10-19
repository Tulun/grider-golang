package main

import (
	"fmt"
	"net/http"
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

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "might be down!")
		c <- "Link might be down"
		return
	}

	fmt.Println(l, "is up!")
	c <- "Link is up"
}
