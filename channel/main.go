package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [5]string{"yea", "hyung", "1", "2", "3"}
	for _, person := range people {
		// go(routines) => non-blocking operation
		// yea, hyung, 1, 2, 3 순서대로 실행되는것을 보장하지 않음
		go isPerson(person, c)
	}
	// <-c => blocking operation
	//fmt.Println("Waiting for messages")
	//fmt.Println("Received message :", <-c)
	//fmt.Println("Received message :", <-c)

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isPerson(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is person"
}
