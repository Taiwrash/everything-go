package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com/",
		"http://facebook.com/",
		"http://amazon.com/",
		"http://twitter.com/",
		"http://reddit.com/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLinks(link, c)
	}

	for l := range c {
		func(l string) {
			time.Sleep(5 * time.Second)
			go checkLinks(l, c)
		}(l)
	}
}

func checkLinks(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, "might be down")
		c <- l
		return
	}

	fmt.Println(l, "is up")
	c <- l
}
