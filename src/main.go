package main

import (
	"log"
	"sample/src/greetings"
	"sample/src/retrybackoff"
)

func main() {
	retrybackoff.TestNewExponentialBackOff()
	message, err := greetings.Hello("Mr.X")

	if err != nil {
		log.Fatal(err)
	}
	log.Println(message)
}
