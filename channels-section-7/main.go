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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(l string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "might be down!")
		return
	}

	fmt.Println(l, "is up!")
}
