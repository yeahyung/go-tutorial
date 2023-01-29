package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan result)
	urls := []string{
		"http://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		tempResult := <-c
		results[tempResult.url] = tempResult.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

//                        send only, cannot receive
func hitURL(url string, c chan<- result) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "ERROR"
	}
	c <- result{
		url: url, status: status,
	}
}
