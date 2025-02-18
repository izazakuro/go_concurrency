package main

import (
	"fmt"
	"os"
	"strings"
)

func printMessage(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping
		if !ok {
			os.Exit(1)
		}

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}

}

func main() {

	ping := make(chan string)
	pong := make(chan string)

	defer close(ping)
	defer close(pong)

	go printMessage(ping, pong)

	fmt.Println("Type something and press Enter(q for quit)")

	for {
		fmt.Print("->")

		var userInput string
		_, _ = fmt.Scanln(&userInput)
		if userInput == "q" {
			break
		}

		ping <- userInput

		responese := <-pong
		fmt.Println("Response: ", responese)
	}

	fmt.Println("Done!")
}
