package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- strings.ToUpper(s)
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	// start gorutine
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER, (enter Q to quit)")

	for {
		fmt.Print("->")
		var userInput string
		_, _ = fmt.Scanln(&userInput)
		if strings.ToLower(userInput) == "q" {
			break
		}
		ping <- userInput
		response := <-pong
		fmt.Println("Response:", response)

	}
	fmt.Println("We are done with the game")
	close(ping)
	close(pong)

}
